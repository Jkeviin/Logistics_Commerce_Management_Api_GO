package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type WarehouseService interface {
	Create(warehouse domain.Warehouse) (domain.Warehouse, error)
	FindAll() ([]domain.Warehouse, error)
	Find(id int64) (domain.Warehouse, error)
	Delete(id int64) error
	Update(id int64, warehouse domain.WarehouseAttributes) (domain.Warehouse, error)
}
