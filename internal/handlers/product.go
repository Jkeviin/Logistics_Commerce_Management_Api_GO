package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
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
func NewProductDefault(sv *services.ProductDefault) *ProductHandler {
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
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/products [get]
func (ph *ProductHandler) GetAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := ph.service.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		// Convertir los productos a ProductDoc
		var productDocs []domain.ProductDoc
		for _, product := range products {
			productDocs = append(productDocs, product.ParseToProductDoc())
		}

		// Devolver los productos en el formato correcto
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
//	@Failure		409		{object}	domain.ErrorResponse		"Código del producto ya existente"
//	@Failure		422		{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500		{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/products [post]
func (ph *ProductHandler) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct domain.ProductDoc

		// Decodificar el cuerpo de la solicitud en una instancia de ProductDoc
		if err := request.JSON(r, &newProduct); err != nil {
			// Devolver un error de solicitud no válida
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrInvalidRequest.Error())
			return
		}
		// Llamar al método ParseToProduct desde la instancia de ProductDoc
		product := newProduct.ParseToProduct()

		productCreated, err := ph.service.Create(product)
		if err != nil {
			if errors.Is(err, utils.ErrProductCodeAlreadyExists) {
				// If the product code already exists, return a conflict response
				response.Error(w, http.StatusConflict, utils.ErrProductCodeAlreadyExists.Error())
				return
			}
			if errors.Is(err, utils.ErrProductAlreadyExists) {
				// if the product already exists, return a conflict response
				response.Error(w, http.StatusConflict, utils.ErrProductAlreadyExists.Error())
				return
			}
			if strings.Contains(err.Error(), "Error al validar el producto") {
				// If the product code already exists, return a conflict response
				response.Error(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			// Devolver un error interno del servidor
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
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
			response.Error(w, http.StatusBadRequest, utils.ErrInvalidProductId.Error())
			return
		}
		product, err := ph.service.GetById(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
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
//	@Failure		422		{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500		{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/products/{id} [patch]
func (ph *ProductHandler) Update() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := utils.ParseIDInt64(idStr)
		if err != nil || id <= 0 {
			response.Error(w, http.StatusBadRequest, utils.ErrInvalidProductId.Error())
			return
		}

		var product domain.ProductAttributesPatch
		if err = request.JSON(r, &product); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrInvalidRequest.Error())
			return
		}

		updatedProduct, err := ph.service.Update(id, product)
		if err != nil {
			if errors.Is(err, utils.ErrProductCodeAlreadyExists) {
				// If the product code already exists, return a conflict response
				response.Error(w, http.StatusConflict, utils.ErrProductCodeAlreadyExists.Error())
				return
			}
			if errors.Is(err, utils.ErrProductAlreadyExists) {
				// if the product already exists, return a conflict response
				response.Error(w, http.StatusConflict, utils.ErrProductAlreadyExists.Error())
				return
			}
			if errors.Is(err, utils.ErrProductNotFound) {
				// If the product is not found, return a not found response
				response.Error(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
				return
			}
			if strings.Contains(err.Error(), "Error al validar el producto") {
				// If the product code already exists, return a conflict response
				response.Error(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			response.Error(w, http.StatusNotFound, utils.ErrInternalServer.Error())
			return
		}

		response.JSON(w, http.StatusOK, domain.ProductResponse{
			Data: updatedProduct.ParseToProductDoc()},
		)
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
//	@Router			/products/{id} [delete]
func (ph *ProductHandler) Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := utils.ParseIDInt64(idStr)
		if err != nil || id <= 0 {
			response.Error(w, http.StatusBadRequest, utils.ErrInvalidProductId.Error())
			return
		}

		err = ph.service.Delete(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
