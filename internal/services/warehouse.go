package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type WarehouseDefault struct {
	rp repository.WarehouseRepository
}

func NewWarehouseDefault(rp repository.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{rp: rp}
}

// Create validates the warehouse data, checks for duplicate warehouse codes,
// and creates a new warehouse if the code is unique.
func (s *WarehouseDefault) Create(warehouse domain.Warehouse) (domain.Warehouse, error) {
	if err := utils.ValidateStructGoValidator(warehouse); err != nil {
		return domain.Warehouse{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return s.rp.Create(warehouse)
}

// FindAll returns all warehouses.
func (s *WarehouseDefault) FindAll() ([]domain.Warehouse, error) {
	return s.rp.FindAll()
}

// Find returns the warehouse with the given ID.
func (s *WarehouseDefault) Find(id int64) (domain.Warehouse, error) {
	return s.rp.Find(id)
}

func (s *WarehouseDefault) Delete(id int64) error {
	return s.rp.Delete(id)
}

func (s *WarehouseDefault) Update(id int64, warehousePatch domain.WarehouseAttributes) (domain.Warehouse, error) {
	// Validate if the warehouse exists and get the warehouse to update
	warehouse, err := s.rp.Find(id)
	if err != nil {
		return domain.Warehouse{}, utils.ErrWareHouseNotFound
	}

	// Update the warehouse with the new data
	utils.UpdateStruct(&warehouse, &warehousePatch)

	// Validate the warehouse data with the govalidator
	if err := utils.ValidateStructGoValidator(warehouse); err != nil {
		return domain.Warehouse{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	// Save the updated warehouse
	return s.rp.Update(id, warehouse)
}
