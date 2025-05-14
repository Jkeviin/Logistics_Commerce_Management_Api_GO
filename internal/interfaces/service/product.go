package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

// ProductService defines the methods for product-related operations.
type ProductService interface {
	GetAll() ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	GetById(id int64) (domain.Product, error)
	Update(id int64, product domain.Product) (domain.Product, error)
	Delete(id int64) error
	FindReports(id *int64) ([]domain.ReportProductRecord, error)
}
