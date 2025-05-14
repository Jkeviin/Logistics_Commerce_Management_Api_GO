package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewProductRecordDefault(rp repository.ProductRecordRepository) *ProductRecordDefault {
	return &ProductRecordDefault{rp: rp}
}

type ProductRecordDefault struct {
	rp repository.ProductRecordRepository
}

func (r *ProductRecordDefault) Create(productRecord domain.ProductRecord) (domain.ProductRecord, error) {
	if err := utils.ValidateStructGoValidator(productRecord); err != nil {
		return domain.ProductRecord{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return r.rp.Create(productRecord)
}
