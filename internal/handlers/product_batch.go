package handlers

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/service"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

func NewProductBatchDefault(sv service.ProductBatchService) *ProductBatchDefault {
	return &ProductBatchDefault{
		sv: sv,
	}
}

type ProductBatchDefault struct {
	sv service.ProductBatchService
}

// Create godoc
//
//	@Summary		Crea un nuevo product batch
//	@Description	Crea un nuevo product batch en el sistema
//	@Tags			ProductBatch
//	@Produce		json
//	@Param			productBatch	body		domain.ProductBatchAttributes	true	"product batch a crear"
//	@Success		201			{object}	domain.ProductBatchResponse	"Product Batch creado exitosamente"
//	@Failure		409			{object}	domain.ErrorResponse		"Conflicto de clave foránea o unica"
//	@Failure		422			{object}	domain.ErrorResponse		"Datos malformados o faltantes"
//	@Failure		500			{object}	domain.ErrorResponse		"Error interno del servidor"
//	@Router			/productBatches [post]
func (p *ProductBatchDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productBatch domain.ProductBatchAttributes
		if err := request.JSON(r, &productBatch); err != nil {
			utils.HandleError(w, utils.ErrBadRequest, nil, http.StatusUnprocessableEntity)
			return
		}

		result, err := p.sv.Create(productBatch.ParseToProductBatch())
		if err != nil {
			utils.HandleError(w, err, map[error]int{
				utils.ErrValidation:        http.StatusUnprocessableEntity,
				utils.ErrProductForeignKey: http.StatusConflict,
				utils.ErrSectionForeignKey: http.StatusConflict,
			}, http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusCreated, domain.ProductBatchResponse{Data: result.ParseToProductBatchDoc()})
	}
}
