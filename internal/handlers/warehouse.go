package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewWarehouseDefault(sv service.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{sv: sv}
}

type WarehouseDefault struct {
	sv service.WarehouseService
}

// Create godoc
//
//	@Summary		Crea un nuevo warehouse
//	@Description	Crea un nuevo warehouse en el sistema
//	@Tags			Warehouse
//	@Produce		json
//	@Param			warehouse	body		domain.WarehouseAttributes	true	"warehouse a crear"
//	@Success		201			{object}	domain.WarehouseResponse	"Warehouse creado exitosamente"
//	@Failure		409			{object}	domain.ErrorResponse		"Código del warehouse ya existente"
//	@Failure		422			{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500			{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/warehouses [post]
func (h *WarehouseDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouse domain.WarehouseAttributes
		if err := request.JSON(r, &warehouse); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Create(warehouse.ParseToWarehouse())
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrWarehouseCodeAlreadyExists: http.StatusConflict,
				utils.ErrValidation:                 http.StatusUnprocessableEntity,
				utils.ErrLocalityForeignKey:         http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.WarehouseResponse{Data: result.ParseToWarehouseDoc()})
	}
}

// GetAll godoc
//
//	@Summary		Lista todos los warehouses
//	@Description	Retorna una lista de todos los warehouses registrados en el sistema
//	@Tags			Warehouse
//	@Produce		json
//	@Success		200	{object}	domain.WarehousesResponse	"Listado de warehouses"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/warehouses [get]
func (h *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAll()
		if err != nil {
			utils.HandleError(w, nil, nil, http.StatusInternalServerError)
			return
		}

		resultDoc := make([]domain.WarehouseDoc, 0, len(result))
		for _, value := range result {
			resultDoc = append(resultDoc, value.ParseToWarehouseDoc())
		}

		response.JSON(w, http.StatusOK, domain.WarehousesResponse{Data: resultDoc})
	}
}

// GetById godoc
//
//	@Summary		Retorna un warehouse por su ID
//	@Description	Retorna la información de un warehouse por su ID
//	@Tags			Warehouse
//	@Produce		json
//	@Param			id	path		int							true	"ID del warehouse"
//	@Success		200	{object}	domain.WarehousesResponse	"Listado de warehouses"
//	@Failure		400	{object}	domain.ErrorResponse		"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse		"El warehouse no existe"
//	@Router			/warehouses/{id} [get]
func (h *WarehouseDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			utils.HandleError(w, err, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Find(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrWareHouseNotFound: http.StatusNotFound,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.WarehouseResponse{Data: result.ParseToWarehouseDoc()})
	}
}

// Delete godoc
//
//	@Summary		Elimina un warehouse por su ID
//	@Description	Elimina un warehouse por su ID de la base de datos
//	@Tags			Warehouse
//	@Produce		json
//	@Param			id	path		int						true	"ID del warehouse"
//	@Success		204	{object}	nil						"warehouse eliminado exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse	"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse	"El warehouse no existe"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/warehouses/{id} [delete]
func (h *WarehouseDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			utils.HandleError(w, err, nil, http.StatusUnprocessableEntity)
			return
		}

		if err := h.sv.Delete(id); err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrWareHouseNotFound:    http.StatusNotFound,
				utils.ErrNoDeleteByForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Update godoc
//
//	@Summary		Actualiza un warehouse por su ID
//	@Description	Actualiza un warehouse por su ID de la base de datos
//	@Tags			Warehouse
//	@Produce		json
//	@Param			id			path		int							true	"ID del warehouse"
//	@Param			warehouse	body		domain.WarehouseAttributes	true	"warehouse a actualizar"
//	@Success		200			{object}	domain.WarehouseResponse	"warehouse actualizado exitosamente"
//	@Failure		400			{object}	domain.ErrorResponse			"ID no enviado o malformado"
//	@Failure		404			{object}	domain.ErrorResponse			"El warehouse no existe"
//	@Failure		409			{object}	domain.ErrorResponse		"Código del warehouse ya existente"
//	@Failure		500			{object}	domain.ErrorResponse			"Error interno del servidor"
//	@Router			/warehouses/{id} [patch]
func (h *WarehouseDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			utils.HandleError(w, err, nil, http.StatusUnprocessableEntity)
			return
		}

		var warehouse domain.WarehouseAttributes
		if err := request.JSON(r, &warehouse); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Update(id, warehouse)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrWarehouseCodeAlreadyExists: http.StatusConflict,
				utils.ErrWareHouseNotFound:          http.StatusNotFound,
				utils.ErrValidation:                 http.StatusUnprocessableEntity,
				utils.ErrLocalityForeignKey:         http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.WarehouseResponse{Data: result.ParseToWarehouseDoc()})
	}
}
