package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

// ProductRepository is an interface that defines the methods for product repository.
type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	FindAllReports() ([]domain.ReportProductRecord, error)
	FindReportById(id int64) (domain.ReportProductRecord, error)
	Create(product domain.Product) (domain.Product, error)
	GetById(id int64) (domain.Product, error)
	Update(product domain.Product) error
	Delete(id int64) error
}
