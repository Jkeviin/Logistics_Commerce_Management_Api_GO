package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type WarehouseRepository interface {
	Create(warehouse domain.Warehouse) (domain.Warehouse, error)
	Find(id int64) (domain.Warehouse, error)
	FindAll() ([]domain.Warehouse, error)
	Delete(id int64) error
	Update(id int64, warehousePointers domain.Warehouse) (domain.Warehouse, error)
}
