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
	// Get all products from the repository
	products, err := p.rp.GetAll()
	if err != nil {
		return nil, err
	}

	//return the list of products
	return products, nil
}

// Create adds a new product to the repository.
func (p *ProductDefault) Create(product domain.Product) (domain.Product, error) {

	//check all fields are not empty or invalid
	if err := utils.ValidateStructGoValidator(product); err != nil {
		return domain.Product{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	// Add the new product to the repository
	productCreated, err := p.rp.Create(product)
	if err != nil {
		return domain.Product{}, err
	}

	return productCreated, nil
}

// GetById returns a product by its ID from the repository.
func (p *ProductDefault) GetById(id int64) (domain.Product, error) {
	// Get the product by ID from the repository
	product, err := p.rp.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// Update updates a product in the repository.
func (p *ProductDefault) Update(id int64, product domain.Product) (domain.Product, error) {
	// Check if the product exists
	if _, err := p.rp.GetById(id); err != nil {
		return domain.Product{}, err
	}

	// Asigna el ID correcto al producto que se va a actualizar
	product.Id = id

	// Validate the product fields
	if err := utils.ValidateStructGoValidator(product); err != nil {
		return domain.Product{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	// Update the product in the repository
	err := p.rp.Update(product)
	if err != nil {
		return domain.Product{}, err
	}

	// Retrieve the updated product from the repository
	updatedProduct, err := p.rp.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return updatedProduct, nil
}

// Delete removes a product from the repository.
func (p *ProductDefault) Delete(id int64) error {
	// Check if the product exists
	_, err := p.rp.GetById(id)
	if err != nil {
		return utils.ErrProductNotFound
	}

	// Delete the product from the repository
	err = p.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductDefault) FindAllReports() ([]domain.ReportProductRecord, error) {
	productRecords, err := r.rp.FindAllReports()
	if err != nil {
		return nil, err
	}
	return productRecords, nil
}

func (r *ProductDefault) FindReports(id *int64) ([]domain.ReportProductRecord, error) {
	if id == nil {
		reports, err := r.rp.FindAllReports()
		return reports, err
	}
	report, err := r.rp.FindReportById(*id)
	if err != nil {
		return nil, err
	}
	return []domain.ReportProductRecord{report}, nil
}
