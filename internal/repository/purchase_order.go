package repository

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type PurchaseOrderSQL struct {
	db *sql.DB
}

func NewPurchaseOrderSQL(db *sql.DB) *PurchaseOrderSQL {
	return &PurchaseOrderSQL{db: db}
}

func (r *PurchaseOrderSQL) Create(po domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	result, err := r.db.Exec(
		queries.CreatePurchaseOrder,
		po.OrderNumber,
		po.OrderDate.Time.Format("2006-01-02"),
		po.TrackingCode,
		po.BuyerID,
		po.ProductRecordID,
	)
	if dbErr := utils.ValidateErrorTypeSQL(err); dbErr != nil {
		if dbErr == utils.ErrDBDuplicateEntry {
			return domain.PurchaseOrder{}, utils.ErrPurchaseOrderConflict
		}
		if dbErr == utils.ErrDBForeignKey {
			if utils.ValidateConstraintFailed(err, "buyer_id") {
				return domain.PurchaseOrder{}, utils.ErrBuyerNotFound
			}
			if utils.ValidateConstraintFailed(err, "product_record_id") {
				return domain.PurchaseOrder{}, utils.ErrProductNotFound
			}
			return domain.PurchaseOrder{}, utils.ErrDBInternalServer
		}
		return domain.PurchaseOrder{}, utils.ErrInternalServer
	}
	if result == nil {
		return domain.PurchaseOrder{}, utils.ErrInternalServer
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.PurchaseOrder{}, utils.ErrInternalServer
	}
	po.ID = id
	return po, nil
}
