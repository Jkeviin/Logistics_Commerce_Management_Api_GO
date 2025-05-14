package services

import (
	"errors"
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewBuyerDefault(rp repository.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{rp: rp}
}

type BuyerDefault struct {
	rp repository.BuyerRepository
}

func (s *BuyerDefault) FindAll() ([]domain.Buyer, error) {
	buyers, err := s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *BuyerDefault) Create(buyer domain.Buyer) (domain.Buyer, error) {
	err := utils.ValidateStructGoValidator(buyer)
	if err != nil {
		return domain.Buyer{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}
	buyer, err = s.rp.Create(buyer)
	if err != nil {
		return domain.Buyer{}, err
	}
	return buyer, nil
}

func (s *BuyerDefault) FindById(id int64) (domain.Buyer, error) {
	buyer, err := s.rp.FindById(id)
	if err != nil {

		return domain.Buyer{}, err
	}
	return buyer, nil
}

func (s *BuyerDefault) Update(id int64, buyerUpdate domain.BuyerAttributes) (buyer domain.Buyer, err error) {
	buyerToUpdate, err := s.rp.FindById(id)
	if err != nil {
		if errors.Is(err, utils.ErrBuyerNotFound) {
			return domain.Buyer{}, utils.ErrBuyerNotFound
		}
		return domain.Buyer{}, utils.ErrInternalServer
	}

	utils.UpdateStruct(&buyerToUpdate, &buyerUpdate)
	err = utils.ValidateStructGoValidator(buyerToUpdate)
	if err != nil {
		return domain.Buyer{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}
	buyerUpdated, err := s.rp.Update(id, buyerToUpdate)
	if err != nil {
		return domain.Buyer{}, err
	}
	return buyerUpdated, nil
}

func (s *BuyerDefault) Delete(id int64) error {
	err := s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BuyerDefault) ReportPurchaseOrders(buyerID *int64) ([]domain.BuyerPurchaseOrdersReport, error) {
	if buyerID != nil {
		return s.rp.ReportPurchaseOrdersByBuyerID(*buyerID)
	}
	return s.rp.ReportPurchaseOrdersAll()
}
