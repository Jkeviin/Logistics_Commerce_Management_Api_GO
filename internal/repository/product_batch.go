package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type ProductBatchSQL struct {
	db *sql.DB
}

func NewProductBatchSQL(db *sql.DB) *ProductBatchSQL {
	return &ProductBatchSQL{db: db}
}

func (p *ProductBatchSQL) Create(productBatch domain.ProductBatch) (domain.ProductBatch, error) {
	result, err := p.db.Exec(
		queries.CreateProductBatch,
		productBatch.BatchNumber,
		productBatch.CurrentQuantity,
		productBatch.CurrentTemperature,
		productBatch.DueDate.Time.Format("2006-01-02"),
		productBatch.InitialQuantity,
		productBatch.ManufacturingDate.Time.Format("2006-01-02"),
		productBatch.ManufacturingHour,
		productBatch.MinimumTemperature,
		productBatch.ProductID,
		productBatch.SectionID,
	)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBForeignKey) {
			if utils.ValidateConstraintFailed(err, "product_id") {
				return domain.ProductBatch{}, utils.ErrProductForeignKey
			}
			if utils.ValidateConstraintFailed(err, "section_id") {
				return domain.ProductBatch{}, utils.ErrSectionForeignKey
			}
		}
		return domain.ProductBatch{}, utils.ErrInternalServer
	}

	if result == nil {
		return domain.ProductBatch{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.ProductBatch{}, utils.ErrInternalServer
	}
	productBatch.ID = id
	return productBatch, nil
}
