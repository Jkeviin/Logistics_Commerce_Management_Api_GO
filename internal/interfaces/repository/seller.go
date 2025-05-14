package repository

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
)

// SellerRepository is an interface that represents a seller repository
type SellerRepository interface {
	// FindAll returns a slice with all sellers
	FindAll() ([]domain.Seller, error)
	// FindByFilter returns a seller that matches the given filter
	FindByFilter(filter filters.SellerFilter) (domain.Seller, error)
	// Create inserts a new seller into the repository
	Create(seller domain.Seller) (domain.Seller, error)
	// Update updates an existing seller in the repository
	Update(id int64, seller domain.Seller) (domain.Seller, error)
	// Delete deletes a seller
	Delete(id int64) error
}
