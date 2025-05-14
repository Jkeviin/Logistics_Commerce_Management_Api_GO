package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

func NewLocalityDefault(sv service.LocalityService) *LocalityDefault {
	return &LocalityDefault{
		sv: sv,
	}
}

type LocalityDefault struct {
	sv service.LocalityService
}

func (h *LocalityDefault) FindSellers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		// Si no se envía id, devolver el reporte de todas las localidades
		if id == "" {
			// Aquí obtienes el reporte para todas las localidades
			result, err := h.sv.FindAllSellers()
			if err != nil {
				utils.HandleError(w, utils.ErrInternalServer, nil, http.StatusInternalServerError)
				return
			}

			response.JSON(w, http.StatusOK, domain.LocalityReportResponse{
				Data: result,
			})
			return
		}

		result, err := h.sv.FindSellers(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrSellersNotFound: http.StatusNotFound,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.LocalityReportResponse{
			Data: []domain.LocalityReportAttributes{result},
		})
	}
}

func (h *LocalityDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var locality domain.LocalityAttributes
		err := request.JSON(r, &locality)
		if err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := h.sv.Create(locality.ParseToLocality())

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrInternalServer: http.StatusInternalServerError,
				utils.ErrValidation:     http.StatusUnprocessableEntity,
				utils.ErrIdAlreadyExist: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.LocalityResponse{
			Data: result.ParseToLocalityDoc(),
		})
	}
}

// GetCantCarriesPerLocality maneja la petición HTTP para obtener la cantidad de carries por localidad.
func (h *LocalityDefault) GetCantCarriesPerLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		result, err := h.sv.GetCantCarriesPerLocality(id)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrInternalServer:     http.StatusInternalServerError,
				utils.ErrLocalityForeignKey: http.StatusNotFound,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, domain.LocalityWithCountResponse{
			Data: result,
		})
	}
}
