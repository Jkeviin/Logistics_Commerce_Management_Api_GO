package repository

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type InboundOrderSQL struct {
	db *sql.DB
}

func NewInboundOrdersSQL(db *sql.DB) *InboundOrderSQL {
	return &InboundOrderSQL{db: db}
}

func (i *InboundOrderSQL) Create(inboundOrder domain.InboundOrder) (domain.InboundOrder, error) {
	res, err := i.db.Exec(
		queries.CreateInboundOrder,
		inboundOrder.OrderDate.Time.Format("2006-01-02"),
		inboundOrder.OrderNumber,
		inboundOrder.Temperature,
		inboundOrder.EmployeeId,
		inboundOrder.ProductBatchId,
		inboundOrder.WarehouseId,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[string]error{
			"uq_order_number":  utils.ErrOrderNumberAlreadyExists,
			"employee_id":      utils.ErrEmployeeNotFound,
			"product_batch_id": utils.ErrProductBatchNotFound,
			"warehouse_id":     utils.ErrWareHouseNotFound,
		}

		for cnt, customErr := range dbErrorsMap {
			if utils.ValidateConstraintFailed(err, cnt) {
				return domain.InboundOrder{}, customErr
			}
		}
		return domain.InboundOrder{}, utils.ErrInternalServer
	}

	if res == nil {
		return domain.InboundOrder{}, utils.ErrInternalServer
	}
	id, err := res.LastInsertId()
	if err != nil {
		return domain.InboundOrder{}, utils.ErrInternalServer
	}
	inboundOrder.ID = id
	return inboundOrder, nil
}
