package handlers

import (
	"net/http"
	"strconv"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
)

func NewEmployeeDefault(sv service.EmployeeService) *EmployeeDefault {
	return &EmployeeDefault{
		sv: sv,
	}
}

type EmployeeDefault struct {
	sv service.EmployeeService
}

// GetAll godoc
//
//	@Summary		Lista todos los Employees
//	@Description	Retorna una lista de todos los Employees registrados en el sistema
//	@Tags			Employee
//	@Produce		json
//	@Success		200	{object}	domain.EmployeesResponse	"Listado de Employees"
//	@Failure		500	{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/employees [get]
func (h *EmployeeDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAll()
		if err != nil {
			utils.HandleError(w, utils.ErrInternalServer, nil, http.StatusInternalServerError)
			return
		}

		// Convertir el slice de empleados a un slice de EmployeeDoc
		employees := make([]domain.EmployeeDoc, 0)
		for _, employee := range result {
			employees = append(employees, employee.ParseToEmployeeDoc())
		}

		// Responder con el JSON correcto
		response.JSON(w, http.StatusOK, domain.EmployeesResponse{
			Data: employees,
		})
	}
}

// GetById godoc
//
//	@Summary		Retorna un Employee por su ID
//	@Description	Retorna la información de un Employee por su ID
//	@Tags			Employee
//	@Produce		json
//	@Param			id	path		int							true	"ID del Employee"
//	@Success		200	{object}	domain.EmployeesResponse	"Listado de Employees"
//	@Failure		400	{object}	domain.ErrorResponse		"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse		"El employee no existe"
//	@Router			/employees/{id} [get]
func (h *EmployeeDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w, utils.ErrMandatoryId, nil, http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			utils.HandleError(w, utils.ErrInvalidId, nil, http.StatusBadRequest)
			return
		}

		result, err := h.sv.Find(idInt)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrEmployeeNotFound: http.StatusNotFound,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.EmployeeResponse{
			Data: result.ParseToEmployeeDoc(),
		})
	}
}

// Create godoc
//
//	@Summary		Crea un nuevo empleado
//	@Description	Crea un nuevo empleado en el sistema con los datos proporcionados
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Param			employee	body		domain.EmployeeAttributes	true	"Datos del empleado a crear"
//	@Success		201			{object}	domain.EmployeeResponse		"Empleado creado exitosamente"
//	@Failure		409			{object}	domain.ErrorResponse		"El código de tarjeta ya existe"
//	@Failure		422			{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500			{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/employees [post]
func (h *EmployeeDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee domain.EmployeeAttributes
		err := request.JSON(r, &employee)
		if err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Create(employee.ParseToEmployee())

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrCardNumberAlreadyExists: http.StatusConflict,
				utils.ErrInternalServer:          http.StatusInternalServerError,
				utils.ErrValidation:              http.StatusUnprocessableEntity,
				utils.ErrWareHouseNotFound:       http.StatusUnprocessableEntity,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.EmployeeResponse{
			Data: result.ParseToEmployeeDoc(),
		})
	}
}

// Delete godoc
//
//	@Summary		Elimina un employee por su ID
//	@Description	Elimina un employee por su ID de la base de datos
//	@Tags			Employee
//	@Produce		json
//	@Param			id	path		int						true	"ID del employee"
//	@Success		204	{object}	nil						"employee eliminado exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse	"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse	"El employee no existe"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/employees/{id} [delete]
func (h *EmployeeDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w, utils.ErrMandatoryId, nil, http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			utils.HandleError(w, utils.ErrInvalidId, nil, http.StatusBadRequest)
			return
		}

		if err := h.sv.Delete(idInt); err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrEmployeeNotFound:     http.StatusNotFound,
				utils.ErrNoDeleteByForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Update godoc
//
//	@Summary		Actualiza un empleado por su ID
//	@Description	Actualiza los datos de un empleado existente en el sistema
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int								true	"ID del empleado"
//	@Param			employee	body		domain.EmployeeAttributes	true	"Datos del empleado a actualizar"
//	@Success		200			{object}	domain.EmployeeResponse			"Empleado actualizado exitosamente"
//	@Failure		400			{object}	domain.ErrorResponse			"ID no enviado o malformado"
//	@Failure		404			{object}	domain.ErrorResponse			"El empleado no existe"
//	@Failure		409			{object}	domain.ErrorResponse			"El código de tarjeta ya existe"
//	@Failure		422			{object}	domain.ErrorResponse			"Datos malformados o faltantes"
//	@Failure		500			{object}	domain.ErrorResponse			"Error interno del servidor"
//	@Router			/employees/{id} [patch]
func (h *EmployeeDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w, utils.ErrMandatoryId, nil, http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			utils.HandleError(w, utils.ErrInvalidId, nil, http.StatusBadRequest)
			return
		}

		var employeePatch domain.EmployeeAttributes
		if err := request.JSON(r, &employeePatch); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Update(idInt, employeePatch)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrCardNumberAlreadyExists: http.StatusConflict,
				utils.ErrInternalServer:          http.StatusInternalServerError,
				utils.ErrEmployeeNotFound:        http.StatusNotFound,
				utils.ErrValidation:              http.StatusUnprocessableEntity,
				utils.ErrWareHouseNotFound:       http.StatusUnprocessableEntity,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.EmployeeResponse{
			Data: result.ParseToEmployeeDoc(),
		})
	}
}

// ReportInboundOrders godoc
//
//	@Summary		Reporte de inbound orders por empleado
//	@Description	Reporte de inbound orders por empleado
//	@Tags			Employee
//	@Produce		json
//	@Param			id	query		int						false	"ID de la sección"
//	@Success		200	{object}	domain.EmployeeWithInboundOrdersCountResponse	"Reporte de inbound orders por empleado"
//	@Failure		400	{object}	domain.ErrorResponse	"ID malformado"
//	@Failure		404	{object}	domain.ErrorResponse	"El empleado no existe"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/employee/reportInboundOrders [get]
func (h *EmployeeDefault) ReportInboundOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		var id int64 = 0
		if idStr != "" {
			var err error
			id, err = strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				utils.HandleError(w, utils.ErrInvalidId, nil, http.StatusBadRequest)
				return
			}
		}

		empl, err := h.sv.ReportInboundOrdersByEmployee(id)

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrEmployeeNotFound: http.StatusNotFound,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.EmployeeWithInboundOrdersCountResponse{
			Data: empl,
		})
	}
}
