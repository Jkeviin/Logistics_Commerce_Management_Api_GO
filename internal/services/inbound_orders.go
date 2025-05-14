package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type InboundOrdersDefault struct {
	rp repository.InboundOrdersRepository
}

func NewInboundOrdersDefault(rp repository.InboundOrdersRepository) *InboundOrdersDefault {
	return &InboundOrdersDefault{rp: rp}
}

// Create validates the product batch data and creates a new inbound order
func (i *InboundOrdersDefault) Create(inboundOrder domain.InboundOrder) (domain.InboundOrder, error) {
	if err := utils.ValidateStructGoValidator(inboundOrder); err != nil {
		return domain.InboundOrder{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}
	return i.rp.Create(inboundOrder)
}
