package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type SectionDefault struct {
	sv service.SectionService
}

func NewSectionDefault(sv service.SectionService) *SectionDefault {
	return &SectionDefault{sv: sv}
}

// Create godoc
//
//	@Summary		Lista todas las sections
//	@Description	Retorna una lista con todas las sections registradas
//	@Tags			Section
//	@Produce		json
//	@Success		200	{object}	domain.SectionsResponse	"Listado de sections"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/sections [get]
func (h *SectionDefault) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		sectionsDoc := make([]domain.SectionDoc, 0, len(sections))
		for _, section := range sections {
			sectionsDoc = append(sectionsDoc, section.ParseToSectionDoc())
		}

		response.JSON(w, http.StatusOK, domain.SectionsResponse{
			Data: sectionsDoc,
		})

	}
}

// FindByID godoc
//
//	@Summary		Retorna un section por su id
//	@Description	Retorna la información de un section por su id
//	@Tags			Section
//	@Produce		json
//	@Param			id	path		int						true	"ID del section"
//	@Success		200	{object}	domain.SectionResponse	"Section encontrado"
//	@Failure		400	{object}	domain.ErrorResponse	"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse	"Section no encontrado"
//	@Router			/sections/{id} [get]
func (h *SectionDefault) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil || id == 0 {
			response.Error(w, http.StatusBadRequest, utils.ErrInvalidId.Error())
			return
		}

		section, err := h.sv.FindById(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, domain.SectionResponse{
			Data: section.ParseToSectionDoc(),
		})

	}
}

// Create godoc
//
//	@Summary		Crea un nuevo section
//	@Description	Crea un nuevo section en el sistema
//	@Tags			Section
//	@Produce		json
//	@Param			section	body		domain.SectionAttributes	true	"section a crear"
//	@Success		201		{object}	domain.SectionResponse		"Section creado exitosamente"
//	@Failure		409		{object}	domain.ErrorResponse		"Código del section ya existente"
//	@Failure		422		{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500		{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/sections [post]
func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var section domain.SectionDoc
		if err := request.JSON(r, &section); err != nil {
			response.Error(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
			return
		}

		err := ValidateSectionPointerAttributes(section.SectionAttributes)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		sectionRes, err := h.sv.Create(section.ParseToSection())

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSectionAlreadyExists: http.StatusConflict,
				utils.ErrInternalServer:       http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}

		response.JSON(w, http.StatusCreated, domain.SectionResponse{
			Data: sectionRes.ParseToSectionDoc(),
		})
	}
}

// Update godoc
//
//	@Summary		Actualiza una sección por su ID
//	@Description	Actualiza una sección por su ID de la base de datos
//	@Tags			Section
//	@Produce		json
//	@Param			id		path		int								true	"ID de la sección"
//	@Param			section	body		domain.SectionAttributes		true	"sección a actualizar"
//	@Success		200		{object}	domain.SectionPatchAttributes	"Sección actualizada exitosamente"
//	@Failure		422		{object}	domain.ErrorResponse			"Datos malformados o faltantes"
//	@Failure		404		{object}	domain.ErrorResponse			"La sección no existe"
//	@Failure		500		{object}	domain.ErrorResponse			"Error interno del servidor"
//	@Router			/sections/{id} [patch]
func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		var sectionPatch domain.SectionPatchAttributes
		if err := request.JSON(r, &sectionPatch); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}

		result, err := h.sv.Update(id, sectionPatch)

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrInternalServer:  http.StatusInternalServerError,
				utils.ErrSectionNotFound: http.StatusNotFound,
			}, http.StatusConflict)
			return
		}

		response.JSON(w, http.StatusOK, domain.SectionResponse{Data: result.ParseToSectionDoc()})
	}
}

// Delete godoc
//
//	@Summary		Elimina una sección por su ID
//	@Description	Elimina una sección por su ID de la base de datos
//	@Tags			Section
//	@Produce		json
//	@Param			id	path		int						true	"ID de la sección"
//	@Success		204	{object}	nil						"sección eliminada exitosamente"
//	@Failure		400	{object}	domain.ErrorResponse	"ID no enviado o malformado"
//	@Failure		404	{object}	domain.ErrorResponse	"La sección no existe"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/sections/{id} [delete]
func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseIDInt64(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.sv.Delete(id); err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSectionNotFound: http.StatusNotFound,
				utils.ErrInternalServer:  http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// AUXILIAR FUNCTIONS
func ValidateSectionPointerAttributes(section domain.SectionAttributes) (err error) {
	// Validar que los campos obligatorios no sean omitidos
	if section.CurrentTemperature == nil || section.MinimumTemperature == nil || section.CurrentCapacity == nil || section.MinimumCapacity == nil {
		err = errors.New("El contenido es inválido o está mal formado")
	}
	return
}
