package handlers

import (
	"errors"
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// NewSellerDefault crea una nueva instancia de SellerDefault
func NewSellerDefault(sv service.SellerService) *SellerDefault {
	return &SellerDefault{
		sv: sv,
	}
}

// SellerDefault implements the handlers for Sellers
type SellerDefault struct {
	// sv is the service that will be used by the handler
	sv service.SellerService
}

// GetAll godoc
//
//	@Summary		Lista todos los proveedores
//	@Description	Retorna una lista de todos los proveedores registrados en el sistema
//	@Tags			Sellers
//	@Produce		json
//	@Success		200	{object}	domain.SellersResponse		"Listado de proveedores"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sellers [get]
func (h *SellerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAll()
		if err != nil {
			if errors.Is(err, utils.ErrSellerNotFound) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		sellersDoc := make([]domain.SellerDoc, len(result))
		for i, seller := range result {
			sellersDoc[i] = seller.ParseToSellerDoc()
		}

		response.JSON(w, http.StatusOK, domain.SellersResponse{
			Data: sellersDoc,
		})
	}
}

// GetById godoc
//
//	@Summary		Retorna un proveedor por su ID
//	@Description	Retorna la información de un proveedor específico basado en su ID
//	@Tags			Sellers
//	@Produce		json
//	@Param			id	path		int	true	"ID del proveedor"
//	@Success		200	{object}	domain.SellerResponse		"Proveedor encontrado"
//	@Failure		400	{object}	domain.ErrorResponse		"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse		"El proveedor no existe"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sellers/{id} [get]
func (h *SellerDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := utils.ParseIDInt64(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		filter := filters.SellerFilter{
			ID: &id,
		}

		result, err := h.sv.FindByFilter(filter)
		if err != nil {
			if errors.Is(err, utils.ErrSellerNotFound) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		response.JSON(w, http.StatusOK, domain.SellerResponse{
			Data: result.ParseToSellerDoc(),
		})
	}
}

// Create godoc
//
//	@Summary		Crea un nuevo proveedor
//	@Description	Crea un nuevo proveedor en el sistema con los datos proporcionados
//	@Tags			Sellers
//	@Accept			json
//	@Produce		json
//	@Param			seller	body		domain.SellerAttributes	true	"Datos del proveedor a crear"
//	@Success		201	{object}	domain.SellerResponse		"Proveedor creado exitosamente"
//	@Failure		409	{object}	domain.ErrorResponse		"El CID ya existe"
//	@Failure		422	{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sellers [post]
func (h *SellerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sellerAttributes domain.SellerAttributes
		if err := request.JSON(r, &sellerAttributes); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}

		seller := sellerAttributes.ParseToSeller()
		result, err := h.sv.Create(seller)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSellerConflict: http.StatusConflict,
				utils.ErrValidation:     http.StatusUnprocessableEntity,
				utils.ErrInternalServer: http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}

		response.JSON(w, http.StatusCreated, domain.SellerResponse{
			Data: result.ParseToSellerDoc(),
		})
	}
}

// Update godoc
//
//	@Summary		Actualiza un proveedor por su ID
//	@Description	Actualiza un proveedor existente con los datos proporcionados
//	@Tags			Sellers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"ID del proveedor"
//	@Param			seller	body		domain.SellerAttributes	true	"Datos del proveedor a actualizar"
//	@Success		200	{object}	domain.SellerResponse		"Proveedor actualizado exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse		"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse		"El proveedor no existe"
//	@Failure		409	{object}	domain.ErrorResponse		"El CID ya existe"
//	@Failure		422	{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sellers/{id} [patch]
func (h *SellerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		var sellerPatch domain.SellerAttributes
		if errRequest := request.JSON(r, &sellerPatch); errRequest != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}

		result, err := h.sv.Update(id, sellerPatch)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSellerNotFound: http.StatusNotFound,
				utils.ErrSellerConflict: http.StatusConflict,
				utils.ErrValidation:     http.StatusUnprocessableEntity,
				utils.ErrInternalServer: http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}

		response.JSON(w, http.StatusOK, domain.SellerResponse{
			Data: result.ParseToSellerDoc(),
		})
	}
}

// Delete godoc
//
//	@Summary		Elimina un proveedor por su ID
//	@Description	Elimina un proveedor existente de la base de datos
//	@Tags			Sellers
//	@Produce		json
//	@Param			id	path		int	true	"ID del proveedor"
//	@Success		204	{object}	nil		"Proveedor eliminado exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse		"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse		"El proveedor no existe"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sellers/{id} [delete]
func (h *SellerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSellerNotFound: http.StatusNotFound,
				utils.ErrInternalServer: http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
