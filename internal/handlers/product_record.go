package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

func NewProductRecordDefault(sv service.ProductRecordService) *ProductRecordDefault {
	return &ProductRecordDefault{
		sv: sv,
	}
}

type ProductRecordDefault struct {
	sv service.ProductRecordService
}

func (h *ProductRecordDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productRecord domain.ProductRecordAttributes
		if err := request.JSON(r, &productRecord); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, utils.ErrBadRequest.Error())
			return
		}

		result, err := h.sv.Create(productRecord.ParseToProductRecord())
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrValidation:        http.StatusUnprocessableEntity,
				utils.ErrProductForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.ProductRecordResponse{Data: result.ParseToProductRecordDoc()})
	}
}
