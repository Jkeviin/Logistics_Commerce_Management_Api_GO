package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type WarehouseMap struct {
	db     map[int64]domain.Warehouse
	loader *loader.JSONLoader[domain.WarehouseDoc, int64]
}

func NewWarehouseMap(db map[int64]domain.Warehouse, loader *loader.JSONLoader[domain.WarehouseDoc, int64]) *WarehouseMap {
	if db == nil {
		db = make(map[int64]domain.Warehouse)
	}
	return &WarehouseMap{db: db, loader: loader}
}

func (r *WarehouseMap) FindAll() ([]domain.Warehouse, error) {
	return utils.MapToSlice(r.db), nil
}

func (r *WarehouseMap) Find(id int64) (domain.Warehouse, error) {
	if warehouse, ok := r.db[id]; ok {
		return warehouse, nil
	}
	return domain.Warehouse{}, utils.ErrWareHouseNotFound
}

func (r *WarehouseMap) Create(warehouse domain.Warehouse) (domain.Warehouse, error) {
	warehouse.ID = time.Now().UnixNano()
	if _, exists := r.db[warehouse.ID]; exists {
		return domain.Warehouse{}, utils.ErrWarehouseAlreadyExists
	}
	r.db[warehouse.ID] = warehouse
	return warehouse, r.saveToJSON()
}

func (r *WarehouseMap) FindByCode(code string) (domain.Warehouse, error) {
	for _, warehouse := range r.db {
		if warehouse.WarehouseCode == code {
			return warehouse, nil
		}
	}
	return domain.Warehouse{}, utils.ErrWareHouseNotFound
}

func (r *WarehouseMap) Delete(id int64) error {
	if _, exists := r.db[id]; !exists {
		return utils.ErrWareHouseNotFound
	}
	delete(r.db, id)
	return r.saveToJSON()
}

func (r *WarehouseMap) Update(id int64, warehouse domain.Warehouse) (domain.Warehouse, error) {
	if _, exists := r.db[id]; !exists {
		return domain.Warehouse{}, utils.ErrWareHouseNotFound
	}
	warehouse.ID = id
	r.db[id] = warehouse
	return warehouse, r.saveToJSON()
}

func (r *WarehouseMap) saveToJSON() error {
	return r.loader.SaveToJSON(r.toWarehouseDoc())
}

func (r *WarehouseMap) toWarehouseDoc() map[int64]domain.WarehouseDoc {
	dbDoc := make(map[int64]domain.WarehouseDoc, len(r.db))
	for id, warehouse := range r.db {
		dbDoc[id] = warehouse.ParseToWarehouseDoc()
	}
	return dbDoc
}
