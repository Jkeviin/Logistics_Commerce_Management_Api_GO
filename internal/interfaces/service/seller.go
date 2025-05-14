package service

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
)

// SellerService is an interface that represents a seller service
type SellerService interface {
	// FindAll returns a slice with all sellers in SellerDoc format
	FindAll() ([]domain.Seller, error)
	// FindByFilter returns a seller that matches the given filter in SellerDoc format
	FindByFilter(filter filters.SellerFilter) (domain.Seller, error)
	// Create creates a new seller and returns it in SellerDoc format
	Create(seller domain.Seller) (domain.Seller, error)
	// Update updates an existing seller and returns it in SellerDoc format
	Update(id int64, seller domain.SellerAttributes) (domain.Seller, error)
	// Delete deletes a seller
	Delete(id int64) error
}
