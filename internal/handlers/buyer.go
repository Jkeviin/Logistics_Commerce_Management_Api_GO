package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
)

func NewBuyerDefault(sv service.BuyerService) *BuyerDefault {
	return &BuyerDefault{
		sv: sv,
	}
}

type BuyerDefault struct {
	sv service.BuyerService
}

// Create godoc
//
//	@Summary		Devuelve todos los buyers
//	@Description	Devuelve todos los buyers del sistema
//	@Tags			Buyer
//	@Produce		json
//	@Success		200	{object}	domain.BuyerListResponse	"Lista de buyers obtenida exitosamente"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/buyers [get]
func (h *BuyerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		buyers, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		buyersDoc := make([]domain.BuyerDoc, len(buyers))
		for i, buyer := range buyers {
			buyersDoc[i] = buyer.ParseToBuyerDoc()
		}

		response.JSON(w, http.StatusOK, domain.BuyerListResponse{
			Data: buyersDoc,
		})
	}
}

// Create godoc
//
//	@Summary		Crea un nuevo buyer
//	@Description	Crea un nuevo buyer en el sistema
//	@Tags			Buyer
//	@Accept			json
//	@Produce		json
//	@Param			buyer	body		domain.BuyerAttributes	true	"Buyer a crear"
//	@Success		201		{object}	domain.BuyerDoc			"Buyer creado exitosamente"
//	@Failure		409		{object}	domain.ErrorResponse	"Código del buyer ya existente"
//	@Failure		422		{object}	domain.ErrorResponse	"Error de validación"
//	@Failure		500		{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/buyers [post]
func (h *BuyerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerAttributes domain.BuyerAttributes
		if err := json.NewDecoder(r.Body).Decode(&buyerAttributes); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		buyer, err := h.sv.Create(buyerAttributes.ParseToBuyer())

		errorMapping := map[error]int{
			utils.ErrBuyerCardNumberAlreadyExists: http.StatusConflict,
			utils.ErrInternalServer:               http.StatusInternalServerError,
			utils.ErrValidation:                   http.StatusUnprocessableEntity,
		}

		if err != nil {
			if status, exists := errorMapping[err]; exists {
				response.Error(w, status, err.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			}
			return
		}

		response.JSON(w, http.StatusCreated, domain.BuyerResponse{
			Data: buyer.ParseToBuyerDoc(),
		})
	}
}

// GetById godoc
//
//	@Summary		Devuelve un buyer por id
//	@Description	Devuelve un buyer por id
//	@Tags			Buyer
//	@Produce		json
//	@Param			id	path		int64					true	"ID del buyer"
//	@Success		200	{object}	domain.BuyerResponse	"Buyer obtenido exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse	"Error de validación"
//	@Failure		404	{object}	domain.ErrorResponse	"Buyer no encontrado"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/buyers/{id} [get]
func (h *BuyerDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		buyer, err := h.sv.FindById(id)
		if err != nil {
			if errors.Is(err, utils.ErrBuyerNotFound) {
				response.Error(w, http.StatusNotFound, utils.ErrBuyerNotFound.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		response.JSON(w, http.StatusOK, domain.BuyerResponse{
			Data: buyer.ParseToBuyerDoc(),
		})
	}
}

// Update godoc
//
//	@Summary		Actualiza un buyer por id
//	@Description	Actualiza un buyer por id
//	@Tags			Buyer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64					true	"ID del buyer"
//	@Param			buyer	body		domain.BuyerAttributes	true	"Buyer a actualizar"
//	@Success		200		{object}	domain.BuyerResponse	"Buyer actualizado exitosamente"
//	@Failure		400		{object}	domain.ErrorResponse	"Error de validación"
//	@Failure		404		{object}	domain.ErrorResponse	"Buyer no encontrado"
//	@Failure		500		{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Failure		409		{object}	domain.ErrorResponse	"Código del buyer ya existente"
//	@Router			/buyers/{id} [patch]
func (h *BuyerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		var buyerAttributes domain.BuyerAttributes
		if err := json.NewDecoder(r.Body).Decode(&buyerAttributes); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		buyer, err := h.sv.Update(id, buyerAttributes)

		errorMapping := map[error]int{
			utils.ErrValidation:     http.StatusConflict,
			utils.ErrBuyerNotFound:  http.StatusNotFound,
			utils.ErrInternalServer: http.StatusInternalServerError,
			utils.ErrBadRequest:     http.StatusBadRequest,
		}

		if err != nil {
			if status, exists := errorMapping[err]; exists {
				response.Error(w, status, err.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, domain.BuyerResponse{
			Data: buyer.ParseToBuyerDoc(),
		})
	}
}

// Delete godoc
//
//	@Summary		Elimina un buyer por id
//	@Description	Elimina un buyer por id
//	@Tags			Buyer
//	@Produce		json
//	@Param			id	path		int64					true	"ID del buyer"
//	@Success		204	{object}	nil						"Buyer eliminado exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse	"Error de validación"
//	@Failure		404	{object}	domain.ErrorResponse	"Buyer no encontrado"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/buyers/{id} [delete]
func (h *BuyerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = h.sv.Delete(id); err != nil {
			if errors.Is(err, utils.ErrBuyerNotFound) {
				response.Error(w, http.StatusNotFound, utils.ErrBuyerNotFound.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)

	}
}
