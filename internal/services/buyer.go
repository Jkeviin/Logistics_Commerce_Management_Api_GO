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
		return domain.Buyer{}, fmt.Errorf("Error al validar el comprador: %w", err)
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
		return domain.Buyer{}, utils.ErrBuyerNotFound
	}
	return buyer, nil
}

func (s *BuyerDefault) Update(id int64, buyerAttr domain.BuyerAttributes) (buyer domain.Buyer, err error) {
	err = utils.ValidateStructGoValidator(buyerAttr)
	if err != nil {
		return domain.Buyer{}, fmt.Errorf("Error al validar el comprador: %w", err)
	}
	buyerToUpdate, err := s.rp.FindById(id)
	if err != nil {
		if errors.Is(err, utils.ErrBuyerNotFound) {
			return domain.Buyer{}, utils.ErrBuyerNotFound
		}
		return domain.Buyer{}, utils.ErrInternalServer
	}
	if buyerAttr.CardNumberId != nil && buyerToUpdate.CardNumberId != *buyerAttr.CardNumberId && s.rp.ExistByCardId(*buyerAttr.CardNumberId) {
		return domain.Buyer{}, utils.ErrValidation
	}
	utils.UpdateStruct(&buyerToUpdate, &buyerAttr)
	buyerUpdated, err := s.rp.Update(id, buyerToUpdate)
	if err != nil {
		if errors.Is(err, utils.ErrBuyerNotFound) {
			return domain.Buyer{}, utils.ErrBuyerNotFound
		}
		return domain.Buyer{}, utils.ErrInternalServer
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
