package handlers

import (
	"errors"
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// ProductHandler handles HTTP requests related to products.
type ProductHandler struct {
	service service.ProductService
}

// NewProductDefault creates a new instance of ProductHandler with the provided service.
func NewProductDefault(sv service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: sv,
	}
}

// GetAll godoc
//
//	@Summary		Lista todos los productos
//	@Description	Retorna una lista de todos los productos registrados en el sistema
//	@Tags			Product
//	@Produce		json
//	@Success		200	{object}	domain.ProductsResponse	"Listado de productos"
//	@Failure		404	{object}	domain.ErrorResponse	"Productos no encontrados"
//	@Failure		400	{object}	domain.ErrorResponse	"Error al parsear la fecha de caducidad"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/products [get]
func (ph *ProductHandler) GetAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := ph.service.GetAll()
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrProductNotFound: http.StatusNotFound,
				utils.ErrFailedParseDate: http.StatusBadRequest,
			}, http.StatusInternalServerError)
			return
		}

		// Convert []domain.Product to []domain.ProductDoc
		productDocs := make([]domain.ProductDoc, len(products))
		for i, product := range products {
			productDocs[i] = product.ParseToProductDoc()
		}

		response.JSON(w, http.StatusOK, domain.ProductsResponse{
			Data: productDocs,
		})
	}
}

// Create godoc
//
//	@Summary		Crea un nuevo producto
//	@Description	Crea un nuevo producto en el sistema
//	@Tags			Product
//	@Produce		json
//	@Param			product	body		domain.ProductAttributes	true	"Producto a crear"
//	@Success		201		{object}	domain.ProductResponse		"Producto creado exitosamente"
//	@Failure		422		{object}	domain.ErrorResponse		"Error de validación"
//	@Failure		409		{object}	domain.ErrorResponse		"Conflicto de producto"
//	@Failure		500		{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Failure		409		{object}	domain.ErrorResponse		"Tipo de producto o vendedor no encontrado"
//	@Router			/products [post]
func (ph *ProductHandler) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct domain.ProductDoc

		// Decodificar el cuerpo de la solicitud en una instancia de ProductDoc
		if err := request.JSON(r, &newProduct); err != nil {
			// Devolver un error de solicitud no válida
			utils.HandleError(w, utils.ErrInvalidRequest, nil, http.StatusUnprocessableEntity)
			return
		}
		// Llamar al metodo ParseToProduct desde la instancia de ProductDoc
		product := newProduct.ParseToProduct()

		productCreated, err := ph.service.Create(product)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrProductCodeAlreadyExists: http.StatusConflict,
				utils.ErrProductAlreadyExists:     http.StatusConflict,
				utils.ErrProductTypeNotFound:      http.StatusConflict,
				utils.ErrSellerNotFound:           http.StatusConflict,
				utils.ErrValidation:               http.StatusUnprocessableEntity,
			}, http.StatusInternalServerError)
			return
		}
		// Return a success response
		response.JSON(w, http.StatusCreated, domain.ProductResponse{
			Data: productCreated.ParseToProductDoc(),
		})
	}
}

// GetByID godoc
//
//	@Summary		Obtiene un producto por su ID
//	@Description	Obtiene un producto por su ID
//	@Tags			Product
//	@Produce		json
//	@Param			id	path		int64					true	"ID del producto"
//	@Success		200	{object}	domain.ProductResponse	"Producto encontrado"
//	@Failure		404	{object}	domain.ErrorResponse	"Producto no encontrado"
//	@Failure		400	{object}	domain.ErrorResponse	"ID de producto inválido"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/products/{id} [get]
func (ph *ProductHandler) GetByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := utils.ParseIDInt64(idStr)
		// Check if the ID is a valid integer
		if err != nil {
			utils.HandleError(w, utils.ErrInvalidProductId, nil, http.StatusBadRequest)
			return
		}
		product, err := ph.service.GetById(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrProductNotFound: http.StatusNotFound,
				utils.ErrFailedParseDate: http.StatusBadRequest,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.ProductResponse{
			Data: product.ParseToProductDoc()})
	}
}

// update patch product
// Update godoc
//
//	@Summary		Actualiza un producto por su ID
//	@Description	Actualiza un producto por su ID
//	@Tags			Product
//	@Produce		json
//	@Param			id		path		int64						true	"ID del producto"
//	@Param			product	body		domain.ProductAttributes	true	"Producto a actualizar"
//	@Success		200		{object}	domain.ProductResponse		"Producto actualizado exitosamente"
//	@Failure		404		{object}	domain.ErrorResponse		"Producto no encontrado"
//	@Failure		400		{object}	domain.ErrorResponse		"ID de producto inválido"
//	@Failure		422		{object}	domain.ErrorResponse		"Error de validación"
//	@Failure		409		{object}	domain.ErrorResponse		"Conflicto de producto"
//	@Failure		500		{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/products/{id} [patch]
func (ph *ProductHandler) Update() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := utils.ParseIDInt64(idStr)
		if err != nil || id <= 0 {
			utils.HandleError(w, utils.ErrInvalidProductId, nil, http.StatusBadRequest)
			return
		}

		var productDoc domain.ProductDoc
		if err := request.JSON(r, &productDoc); err != nil {
			utils.HandleError(w, utils.ErrInvalidRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		updatedProduct, err := ph.service.Update(id, productDoc.ParseToProduct())
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrProductNotFound:          http.StatusNotFound,
				utils.ErrProductCodeAlreadyExists: http.StatusConflict,
				utils.ErrProductTypeNotFound:      http.StatusConflict,
				utils.ErrSellerNotFound:           http.StatusConflict,
				utils.ErrValidation:               http.StatusUnprocessableEntity,
				utils.ErrFailedParseDate:          http.StatusBadRequest,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.ProductResponse{
			Data: updatedProduct.ParseToProductDoc(),
		})
	}
}

// Delete godoc
//
//	@Summary		Elimina un producto por su ID
//	@Description	Elimina un producto por su ID
//	@Tags			Product
//	@Produce		json
//	@Param			id	path		int64					true	"ID del producto"
//	@Success		204	{object}	domain.ProductResponse	"Producto eliminado exitosamente"
//	@Failure		404	{object}	domain.ErrorResponse	"Producto no encontrado"
//	@Failure		400	{object}	domain.ErrorResponse	"ID de producto inválido"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Failure		409	{object}	domain.ErrorResponse	"Error de eliminación por clave foránea"
//	@Router			/products/{id} [delete]
func (ph *ProductHandler) Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil || id <= 0 {
			utils.HandleError(w, utils.ErrInvalidProductId, nil, http.StatusBadRequest)
			return
		}

		err = ph.service.Delete(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrProductNotFound:      http.StatusNotFound,
				utils.ErrNoDeleteByForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

// GetAll godoc
//
//	@Summary		Devuelve todos los reportes de productos y filtra por id de producto
//	@Description	Devuelve todos los reportes de productos y filtra por id de producto
//	@Tags			ProductRecord
//	@Produce		json
//
// @Param			id	query		int	false	"ID del producto"
//
//	@Success		200	{object}	domain.ReportResponse	"Lista de reportes obtenida exitosamente"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/productRecords [get]
func (h *ProductHandler) GetReportRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productId *int64
		if idStr := r.URL.Query().Get("id"); idStr != "" {
			id, err := utils.ParseIDInt64(idStr)
			if err != nil {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			productId = &id
		}
		productRecords, err := h.service.FindReports(productId)
		if err != nil {
			if errors.Is(err, utils.ErrProductNotFound) {
				response.Error(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		response.JSON(w, http.StatusOK, domain.ReportResponse{
			Data: productRecords,
		})
	}
}
