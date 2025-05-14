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
		return domain.Warehouse{}, fmt.Errorf("Error al validar el almacén: %w", err)
	}

	if _, err := s.rp.FindByCode(warehouse.WarehouseCode); err == nil {
		return domain.Warehouse{}, utils.ErrWarehouseCodeAlreadyExists
	}

	result, err := s.rp.Create(warehouse)
	if err != nil {
		return domain.Warehouse{}, utils.ErrInternalServer
	}
	return result, nil
}

// FindAll returns all warehouses.
func (s *WarehouseDefault) FindAll() ([]domain.Warehouse, error) {
	result, err := s.rp.FindAll()
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	return result, nil
}

// Find returns the warehouse with the given ID.
func (s *WarehouseDefault) Find(id int64) (domain.Warehouse, error) {
	result, err := s.rp.Find(id)
	if err != nil {
		return domain.Warehouse{}, utils.ErrWareHouseNotFound
	}
	return result, nil
}

func (s *WarehouseDefault) Delete(id int64) error {
	if _, err := s.rp.Find(id); err != nil {
		return utils.ErrWareHouseNotFound
	}
	if err := s.rp.Delete(id); err != nil {
		return utils.ErrInternalServer
	}
	return nil
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
		return domain.Warehouse{}, fmt.Errorf("Error al validar el almacén: %w", err)
	}

	// Check if the warehouse code already exists
	if existing, err := s.rp.FindByCode(warehouse.WarehouseCode); err == nil && existing.ID != id {
		return domain.Warehouse{}, utils.ErrWarehouseCodeAlreadyExists
	}

	// Save the updated warehouse
	result, err := s.rp.Update(id, warehouse)
	if err != nil {
		return domain.Warehouse{}, utils.ErrInternalServer
	}
	return result, nil
}
