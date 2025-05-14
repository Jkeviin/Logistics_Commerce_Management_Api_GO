package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewProductRecordSQL(db *sql.DB) *ProductRecordSQL {
	return &ProductRecordSQL{
		db: db,
	}
}

type ProductRecordSQL struct {
	db *sql.DB
}

func (r *ProductRecordSQL) Create(productRecord domain.ProductRecord) (domain.ProductRecord, error) {
	result, err := r.db.Exec(
		queries.CreateProductRecord,
		productRecord.LastUpdateDate.Time.Format("2006-01-02"),
		productRecord.PurchasePrice,
		productRecord.SalePrice,
		productRecord.ProductId,
	)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBForeignKey) {
			return domain.ProductRecord{}, utils.ErrProductForeignKey
		}
		return domain.ProductRecord{}, utils.ErrInternalServer
	}

	if result == nil {
		return domain.ProductRecord{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.ProductRecord{}, utils.ErrInternalServer
	}
	productRecord.Id = id
	return productRecord, nil
}
