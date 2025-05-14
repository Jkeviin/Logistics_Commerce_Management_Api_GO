package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

// ProductRepository is an interface that defines the methods for product repository.
type ProductRepository interface {
	GetAll() (map[int64]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	GetById(id int64) (domain.Product, error)
	Update(product domain.Product) error
	GetByCode(code string) (domain.Product, error)
	Delete(id int64) error
}
