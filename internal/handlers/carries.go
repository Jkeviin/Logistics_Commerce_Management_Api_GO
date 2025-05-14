package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

// ProductHandler handles HTTP requests related to products.
type CarriesHandler struct {
	service service.CarriesService
}

// NewProductDefault creates a new instance of ProductHandler with the provided service.
func NewCarriesDefault(sv *services.CarriesDefault) *CarriesHandler {
	return &CarriesHandler{
		service: sv,
	}
}

// Create godoc
//
//	@Summary		Crea un nuevo carry
//	@Description	Crea un nuevo carry en el sistema
//	@Tags			Carries
//	@Accept			json
//	@Produce		json
//	@Param			carries	body		domain.CarriesDoc	true	"Datos del carry"
//	@Success		201	{object}	domain.CarryResponse	"Carry creado"
//	@Failure		422	{object}	domain.ErrorResponse	"Error de validación"
//	@Failure		409	{object}	domain.ErrorResponse	"Carry ya existe"
//	@Failure		500	{object}	domain.ErrorResponse	"Error interno del servidor"
//	@Router			/carries [post]
func (ph *CarriesHandler) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCarries domain.CarriesDoc

		// Decodificar el cuerpo de la solicitud en una instancia de Carries
		if err := request.JSON(r, &newCarries); err != nil {
			// Devolver un error de solicitud no válida
			utils.HandleError(w, utils.ErrInvalidRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		carries := newCarries.ParseToCarrie()

		carrieCreated, err := ph.service.Create(carries)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrCarriesAlreadyExists: http.StatusConflict,
				utils.ErrLocalityForeignKey:   http.StatusConflict,
				utils.ErrValidation:           http.StatusUnprocessableEntity,
			}, http.StatusInternalServerError)
			return
		}
		response.JSON(w, http.StatusCreated, domain.CarryResponse{
			Data: carrieCreated.ParseToCarrieDoc(),
		})
	}
}
