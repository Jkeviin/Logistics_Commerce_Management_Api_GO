package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type ProductBatchDefault struct {
	rp repository.ProductBatchRepository
}

func NewProductBatchDefault(rp repository.ProductBatchRepository) *ProductBatchDefault {
	return &ProductBatchDefault{rp: rp}
}

// Create validates the product batch data and creates a new product batch.
func (p *ProductBatchDefault) Create(productBatch domain.ProductBatch) (domain.ProductBatch, error) {
	if err := utils.ValidateStructGoValidator(productBatch); err != nil {
		return domain.ProductBatch{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return p.rp.Create(productBatch)
}
