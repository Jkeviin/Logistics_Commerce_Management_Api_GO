package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// ProductRepository is an interface that defines the methods for product repository.
func NewProductMap(db map[int64]domain.Product, loader *loader.JSONLoader[domain.Product, int64]) *ProductMap {
	// default db
	defaultDb := make(map[int64]domain.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{
		db:     defaultDb,
		loader: loader,
	}
}

// ProductMap is a struct that implements the ProductRepository interface.
type ProductMap struct {
	// db is a map of products
	db map[int64]domain.Product
	// loader is a JSON loader
	loader *loader.JSONLoader[domain.Product, int64]
}

// GetAll returns all products from the repository.
func (p *ProductMap) GetAll() (map[int64]domain.Product, error) {
	// return all products
	return p.db, nil
}

// Create adds a new product to the repository.
func (p *ProductMap) Create(product domain.Product) (domain.Product, error) {
	// add product to db
	product.Id = time.Now().UnixNano()
	p.db[product.Id] = product

	// load product
	if err := p.loader.SaveToJSON(p.db); err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}
	return product, nil
}

// GetById returns a product by its ID from the repository.
func (p *ProductMap) GetById(id int64) (domain.Product, error) {
	// get product by id
	product, ok := p.db[id]
	if !ok {
		return domain.Product{}, utils.ErrProductNotFound
	}
	return product, nil
}

// update updates a product in the repository.
func (p *ProductMap) Update(product domain.Product) error {
	// check if product exists
	_, ok := p.db[product.Id]
	if !ok {
		return utils.ErrProductNotFound
	}
	// update product
	p.db[product.Id] = product

	// load product
	if err := p.loader.SaveToJSON(p.db); err != nil {
		return utils.ErrProductNotFound
	}
	return nil
}

// find by code
func (p *ProductMap) GetByCode(code string) (domain.Product, error) {
	// find product by code
	for _, product := range p.db {
		if product.ProductCode == code {
			return domain.Product{}, utils.ErrProductCodeAlreadyExists
		}
	}
	return domain.Product{}, nil
}

// Delete removes a product from the repository.
func (p *ProductMap) Delete(id int64) error {
	// check if product exists
	_, ok := p.db[id]
	if !ok {
		return utils.ErrProductNotFound
	}
	// delete product
	delete(p.db, id)

	// load product
	if err := p.loader.SaveToJSON(p.db); err != nil {
		return utils.ErrInternalServer
	}
	return nil
}
