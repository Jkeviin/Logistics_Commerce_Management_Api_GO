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
		var buyerCreate domain.BuyerAttributes
		if err := json.NewDecoder(r.Body).Decode(&buyerCreate); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}
		buyer, err := h.sv.Create(buyerCreate.ParseToBuyer())
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrBuyerCardNumberAlreadyExists: http.StatusConflict,
				utils.ErrValidation:                   http.StatusUnprocessableEntity,
			}, http.StatusInternalServerError)
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

		var buyerUpdate domain.BuyerAttributes
		if err := json.NewDecoder(r.Body).Decode(&buyerUpdate); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}

		buyer, err := h.sv.Update(id, buyerUpdate)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrBuyerNotFound:                http.StatusNotFound,
				utils.ErrBadRequest:                   http.StatusBadRequest,
				utils.ErrValidation:                   http.StatusUnprocessableEntity,
				utils.ErrBuyerCardNumberAlreadyExists: http.StatusConflict,
			}, http.StatusInternalServerError)
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
			utils.HandleError(w, err, map[error]int{
				utils.ErrBuyerNotFound:        http.StatusNotFound,
				utils.ErrNoDeleteByForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)

	}
}

// ReportPurchaseOrders godoc
//
//	@Summary		Reporte de órdenes de compra por Buyer
//	@Description	Devuelve un reporte de órdenes de compra por Buyer
//	@Tags			Buyer
//	@Produce		json
//	@Param			id	query		int64					false	"ID del buyer"
//	@Success		200		{object}	domain.BuyerPurchaseOrdersReportResponse	"Reporte de órdenes de compra obtenido exitosamente"
//	@Failure		400		{object}	domain.ErrorResponse						"Error de validación"
//	@Failure		404		{object}	domain.ErrorResponse						"Buyer no encontrado"
//	@Failure		500		{object}	domain.ErrorResponse						"Error interno del servidor"
//	@Router			/buyers/reportPurchaseOrders [get]
func (h *BuyerDefault) ReportPurchaseOrders() http.HandlerFunc {
	// Distinct error maps for clarity
	errorMapParse := map[error]int{
		utils.ErrBadRequest: http.StatusBadRequest,
	}
	errorMapDomain := map[error]int{
		utils.ErrBuyerNotFound:  http.StatusNotFound,
		utils.ErrInternalServer: http.StatusInternalServerError,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerID *int64
		if idStr := r.URL.Query().Get("id"); idStr != "" {
			id, err := utils.ParseIDInt64(idStr)
			if err != nil {
				utils.HandleError(w, utils.ErrBadRequest, errorMapParse, http.StatusBadRequest)
				return
			}
			buyerID = &id
		}
		reports, err := h.sv.ReportPurchaseOrders(buyerID)
		if err != nil {
			utils.HandleError(w, err, errorMapDomain, http.StatusInternalServerError)
			return
		}
		response.JSON(w, http.StatusOK, domain.BuyerPurchaseOrdersReportResponse{Data: reports})
	}
}
