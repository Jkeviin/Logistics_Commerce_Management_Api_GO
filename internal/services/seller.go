package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// NewSellerDefault creates a new instance of SellerDefault
func NewSellerDefault(rp repository.SellerRepository) *SellerDefault {
	return &SellerDefault{rp: rp}
}

// SellerDefault is a struct that represents the default service for sellers
type SellerDefault struct {
	// rp is the repository that will be used by the service
	rp repository.SellerRepository
}

// FindAll returns a slice of all sellers in Seller format
func (s *SellerDefault) FindAll() ([]domain.Seller, error) {
	return s.rp.FindAll()
}

// FindByFilter returns a seller that matches the given filter in Seller format
func (s *SellerDefault) FindByFilter(filter filters.SellerFilter) (domain.Seller, error) {
	return s.rp.FindByFilter(filter)
}

// Create creates a new seller and returns it in Seller format
func (s *SellerDefault) Create(seller domain.Seller) (domain.Seller, error) {
	if err := utils.ValidateStructGoValidator(seller); err != nil {
		return domain.Seller{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}
	return s.rp.Create(seller)
}

// Update updates an existing seller and returns it in Seller format
func (s *SellerDefault) Update(id int64, sellerPatch domain.SellerAttributes) (domain.Seller, error) {
	existingSeller, err := s.rp.FindByFilter(filters.SellerFilter{ID: &id})
	if err != nil {
		return domain.Seller{}, err
	}

	utils.UpdateStruct(&existingSeller, &sellerPatch)

	if err := utils.ValidateStructGoValidator(existingSeller); err != nil {
		return domain.Seller{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return s.rp.Update(id, existingSeller)
}

// Delete deletes a seller
func (s *SellerDefault) Delete(id int64) error {
	return s.rp.Delete(id)
}
