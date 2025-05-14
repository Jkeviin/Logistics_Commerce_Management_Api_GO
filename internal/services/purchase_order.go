package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type PurchaseOrderDefault struct {
	poRepo repository.PurchaseOrderRepository
}

func NewPurchaseOrderDefault(poRepo repository.PurchaseOrderRepository) *PurchaseOrderDefault {
	return &PurchaseOrderDefault{
		poRepo: poRepo,
	}
}

func (s *PurchaseOrderDefault) Create(po domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	if err := utils.ValidateStructGoValidator(po); err != nil {
		return domain.PurchaseOrder{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	created, err := s.poRepo.Create(po)
	if err != nil {
		return domain.PurchaseOrder{}, err
	}
	return created, nil
}
