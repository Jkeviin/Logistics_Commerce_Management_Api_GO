package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type PurchaseOrderDefault struct {
	sv service.PurchaseOrderService
}

func NewPurchaseOrderDefault(sv service.PurchaseOrderService) *PurchaseOrderDefault {
	return &PurchaseOrderDefault{sv: sv}
}

// Create godoc
//
// @Summary	Crea una nueva orden de compra
// @Description	Crea una nueva orden de compra en el sistema
// @Tags	PurchaseOrders
// @Accept	json
// @Produce	json
// @Param	purchase_order	body	domain.PurchaseOrderAttributesSwagger	true	"Datos de la orden de compra a crear"
// @Success	201	{object}	domain.PurchaseOrderResponse	"Orden de compra creada exitosamente"
// @Failure	409	{object}	domain.ErrorResponse	"El número de orden ya existe, o buyer/product no existen"
// @Failure	422	{object}	domain.ErrorResponse	"Datos malformados o faltantes"
// @Failure	500	{object}	domain.ErrorResponse	"Error interno del servidor"
// @Router	/purchaseOrders [post]
func (h *PurchaseOrderDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var attrs domain.PurchaseOrderAttributes
		if err := request.JSON(r, &attrs); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, map[error]int{
				utils.ErrBadRequest: http.StatusBadRequest,
			}, http.StatusBadRequest)
			return
		}
		po, parseErr := attrs.ParseToPurchaseOrder()
		if parseErr != nil {
			utils.HandleError(w, utils.ErrBadRequest, map[error]int{
				utils.ErrBadRequest: http.StatusBadRequest,
			}, http.StatusBadRequest)
			return
		}
		po, err := h.sv.Create(po)
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				// Validation errors
				utils.ErrValidation: http.StatusUnprocessableEntity,
				// Domain/business errors
				utils.ErrPurchaseOrderConflict: http.StatusConflict,
				utils.ErrBuyerNotFound:         http.StatusConflict,
				utils.ErrProductNotFound:       http.StatusConflict,
				utils.ErrInternalServer:        http.StatusInternalServerError,
			}, http.StatusUnprocessableEntity)
			return
		}
		response.JSON(w, http.StatusCreated, domain.PurchaseOrderResponse{
			Data: po.ParseToPurchaseOrderDoc(),
		})
	}
}
