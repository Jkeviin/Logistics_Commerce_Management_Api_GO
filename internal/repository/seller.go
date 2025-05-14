package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// NewSellerMap creates a new instance of SellerMap
func NewSellerMap(db map[int64]domain.Seller, loader *loader.JSONLoader[domain.SellerDoc, int64]) repository.SellerRepository {
	return &SellerMap{
		db:     db,
		loader: loader,
	}
}

// SellerMap is a structure that represents a seller repository
type SellerMap struct {
	// db is a map of sellers
	db     map[int64]domain.Seller
	loader *loader.JSONLoader[domain.SellerDoc, int64]
}

// FindAll returns a slice with all sellers
func (r *SellerMap) FindAll() ([]domain.Seller, error) {
	return utils.MapToSlice(r.db), nil
}

// FindByFilter returns a seller that matches the given filter
func (r *SellerMap) FindByFilter(filter filters.SellerFilter) (domain.Seller, error) {
	if !filter.HasFilters() {
		return domain.Seller{}, utils.ErrInvalidFilter
	}

	// Si el filtro incluye ID, buscar directamente por ID
	if filter.ID != nil && *filter.ID > 0 {
		seller, exists := r.db[*filter.ID]
		if !exists {
			return domain.Seller{}, utils.ErrSellerNotFound
		}
		return seller, nil
	}

	// Para otros filtros, buscar en el mapa
	for _, seller := range r.db {
		if filter.Matches(seller) {
			return seller, nil
		}
	}

	return domain.Seller{}, utils.ErrSellerNotFound
}

// Create inserts a new seller into the repository
func (r *SellerMap) Create(seller domain.Seller) (domain.Seller, error) {
	// Check if CID already exists using filter
	filter := filters.SellerFilter{
		CID: &seller.CID,
	}

	for _, s := range r.db {
		if filter.Matches(s) {
			return domain.Seller{}, utils.ErrSellerConflict
		}
	}

	// Generate new ID using UnixNano
	seller.ID = time.Now().UnixNano()

	// Add to map
	r.db[seller.ID] = seller

	// Save to JSON
	if err := r.loader.SaveToJSON(r.toSellerDoc()); err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

// Update updates an existing seller in the repository
func (r *SellerMap) Update(id int64, seller domain.Seller) (domain.Seller, error) {
	// Check if seller exists
	_, exists := r.db[id]
	if !exists {
		return domain.Seller{}, utils.ErrSellerNotFound
	}

	// Update seller in map
	r.db[id] = seller

	// Save to JSON
	if err := r.loader.SaveToJSON(r.toSellerDoc()); err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

// Delete deletes a seller from the repository
func (r *SellerMap) Delete(id int64) error {
	// Check if seller exists
	seller, exists := r.db[id]
	if !exists {
		return utils.ErrSellerNotFound
	}

	// Delete seller from map
	delete(r.db, id)

	// Save to JSON
	if err := r.loader.SaveToJSON(r.toSellerDoc()); err != nil {
		// Rollback: restore the deleted seller
		r.db[id] = seller
		return err
	}

	return nil
}

// AUXILIAR FUNCTIONS
func (r *SellerMap) toSellerDoc() map[int64]domain.SellerDoc {
	dbDoc := make(map[int64]domain.SellerDoc)
	for key, value := range r.db {
		dbDoc[key] = value.ParseToSellerDoc()
	}
	return dbDoc
}
