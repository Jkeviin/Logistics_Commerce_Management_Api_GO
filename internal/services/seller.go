package services

import (
	"errors"
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
	seller, err := s.rp.FindByFilter(filter)
	if err != nil {
		if errors.Is(err, utils.ErrSellerNotFound) {
			return domain.Seller{}, err
		}
		return domain.Seller{}, utils.ErrInternalServer
	}
	return seller, nil
}

// Create creates a new seller and returns it in Seller format
func (s *SellerDefault) Create(seller domain.Seller) (domain.Seller, error) {
	// Validate seller struct
	if err := utils.ValidateStructGoValidator(seller); err != nil {
		return domain.Seller{}, fmt.Errorf("Error al validar el Proveedor: %w", err)
	}

	createdSeller, err := s.rp.Create(seller)
	if err != nil {
		if errors.Is(err, utils.ErrSellerConflict) {
			return domain.Seller{}, err
		}
		return domain.Seller{}, utils.ErrInternalServer
	}
	return createdSeller, nil
}

// Update updates an existing seller and returns it in Seller format
func (s *SellerDefault) Update(id int64, sellerPatch domain.SellerAttributes) (domain.Seller, error) {
	// Get existing seller
	existingSeller, errFind := s.rp.FindByFilter(filters.SellerFilter{ID: &id})
	if errFind != nil {
		if errors.Is(errFind, utils.ErrSellerNotFound) {
			return domain.Seller{}, utils.ErrSellerNotFound
		}
		return domain.Seller{}, utils.ErrInternalServer
	}

	// Check if CID is being updated and if it already exists
	if sellerPatch.CID != nil && *sellerPatch.CID != existingSeller.CID {
		filter := filters.SellerFilter{
			CID: sellerPatch.CID,
		}
		conflictingSeller, err := s.rp.FindByFilter(filter)
		if err == nil && conflictingSeller.ID != id {
			return domain.Seller{}, utils.ErrSellerConflict
		}
	}

	utils.UpdateStruct(&existingSeller, &sellerPatch)

	// Validate
	if err := utils.ValidateStructGoValidator(existingSeller); err != nil {
		return domain.Seller{}, fmt.Errorf("Error al validar el Proveedor: %w", err)
	}

	// Update the seller
	updatedSeller, err := s.rp.Update(id, existingSeller)
	if err != nil {
		if errors.Is(err, utils.ErrSellerNotFound) {
			return domain.Seller{}, err
		}
		return domain.Seller{}, utils.ErrInternalServer
	}
	return updatedSeller, nil
}

// Delete deletes a seller
func (s *SellerDefault) Delete(id int64) error {
	err := s.rp.Delete(id)
	if err != nil {
		if errors.Is(err, utils.ErrSellerNotFound) {
			return err
		}
		return utils.ErrInternalServer
	}
	return nil
}
