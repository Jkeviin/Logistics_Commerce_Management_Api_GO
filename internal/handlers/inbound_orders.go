package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

func NewInboundOrdersDefault(sv service.InboundOrdersService) *InboundOrdersDefault {
	return &InboundOrdersDefault{
		sv: sv,
	}
}

type InboundOrdersDefault struct {
	sv service.InboundOrdersService
}

// Create godoc
//
//	@Summary		Crea un nuevo inbound order
//	@Description	Crea un nuevo inbound en el sistema
//	@Tags			InboundOrder
//	@Produce		json
//	@Param			InboundOrder	body		domain.InboundOrderAttributes	true	"inbound order a crear"
//	@Success		201			{object}	domain.InboundOrderResponse	"Inbound Order creado exitosamente"
//	@Failure		409			{object}	domain.ErrorResponse		"Conflicto de clave foránea o unica"
//	@Failure		422			{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500			{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/inboundOrders [post]
func (i *InboundOrdersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inboundOrder domain.InboundOrderAttributes
		if err := request.JSON(r, &inboundOrder); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := i.sv.Create(inboundOrder.ParseToInboundOrder())

		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrInternalServer:           http.StatusInternalServerError,
				utils.ErrValidation:               http.StatusUnprocessableEntity,
				utils.ErrOrderNumberAlreadyExists: http.StatusConflict,
				utils.ErrEmployeeNotFound:         http.StatusConflict,
				utils.ErrProductBatchNotFound:     http.StatusConflict,
				utils.ErrWareHouseNotFound:        http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.InboundOrderResponse{Data: result.ParseToInboundOrderDoc()})
	}
}
