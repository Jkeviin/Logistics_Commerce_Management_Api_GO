package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// ProductService is an interface that defines the methods for product service.
func NewProductDefault(rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

// ProductDefault is a struct that implements the ProductService interface.
type ProductDefault struct {
	rp repository.ProductRepository
}

// GetAll returns all products from the repository.
func (p *ProductDefault) GetAll() ([]domain.Product, error) {
	// Obtain all products from the repository
	products, err := p.rp.GetAll()
	// Check if there was an error or if no products were found
	if err != nil {
		return nil, utils.ErrProductNotFound
	}

	// Manually convert the map of products to a slice
	productSlice := utils.MapToSlice(products)
	return productSlice, nil
}

// Create adds a new product to the repository.
func (p *ProductDefault) Create(product domain.Product) (domain.Product, error) {

	//check all fields are not empty or invalid
	if err := utils.ValidateStructGoValidator(product); err != nil {
		return domain.Product{}, fmt.Errorf("Error al validar el producto: %w", err)
	}

	// Check if the product code already exists
	_, err := p.rp.GetByCode(product.ProductCode)
	if err != nil {
		return domain.Product{}, utils.ErrProductCodeAlreadyExists
	}

	// Add the new product to the repository
	productCreated, err := p.rp.Create(product)
	if err != nil {
		return domain.Product{}, utils.ErrProductAlreadyExists
	}

	return productCreated, nil
}

// GetById returns a product by its ID from the repository.
func (p *ProductDefault) GetById(id int64) (domain.Product, error) {
	// Get the product by ID from the repository
	product, err := p.rp.GetById(id)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}
	return product, nil
}

// Update updates a product in the repository.
func (p *ProductDefault) Update(id int64, product domain.ProductAttributesPatch) (domain.Product, error) {
	// Obtener el producto existente por ID
	existingProduct, err := p.rp.GetById(id)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}
	// Actualizar los atributos del producto existente con los nuevos valores

	utils.UpdateStruct(&existingProduct, &product)

	if err := utils.ValidateStructGoValidator(existingProduct); err != nil {
		return domain.Product{}, fmt.Errorf("Error al validar el producto: %s", err)
	}

	_, err = p.rp.GetByCode(*product.ProductCode)
	if err != nil {
		return domain.Product{}, utils.ErrProductCodeAlreadyExists
	}

	//check if the product id already exists
	productSearched, err := p.rp.GetById(id)
	if err == nil && productSearched.Id != id {
		return domain.Product{}, utils.ErrProductCodeAlreadyExists
	}
	//save the updated product
	err = p.rp.Update(existingProduct)
	if err != nil {
		return domain.Product{}, utils.ErrProductNotFound
	}
	return existingProduct, nil
}

// FindByCode returns a product by its code from the repository.
func (p *ProductDefault) GetByCode(code string) (domain.Product, error) {
	// Get the product by code from the repository
	product, err := p.rp.GetByCode(code)
	if err != nil {
		return domain.Product{}, utils.ErrProductCodeAlreadyExists
	}
	return product, nil
}

// Delete removes a product from the repository.
func (p *ProductDefault) Delete(id int64) error {
	// Check if the product exists
	product, err := p.rp.GetById(id)
	if err != nil {
		return utils.ErrProductNotFound
	}

	// Delete the product from the repository
	err = p.rp.Delete(product.Id)
	if err != nil {
		return utils.ErrInternalServer
	}
	return nil
}
