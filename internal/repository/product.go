package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// ProductRepository is an interface that defines the methods for product repository.
func NewProductSQL(db *sql.DB) *ProductSQL {
	return &ProductSQL{
		db: db,
	}
}

// ProductSQL is a struct that implements the ProductRepository interface.
type ProductSQL struct {
	db *sql.DB
}

// GetAll returns all products from the repository.
func (p *ProductSQL) GetAll() ([]domain.Product, error) {
	// get all products from db
	rows, err := p.db.Query(queries.GetAllProducts)
	if err != nil {
		return nil, err
	}
	defer utils.CloseRows(rows)

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		var expirationDate []uint8
		if err := rows.Scan(&product.Id, &product.ProductCode, &product.Description, &product.Width, &product.Height,
			&product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature,
			&product.FreezingRate, &product.ProductTypeId, &product.SellerId, &expirationDate); err != nil {
			return nil, err
		}
		parsedDate, err := utils.ParseDateTime(expirationDate)
		if err != nil {
			return nil, utils.ErrFailedParseDate
		}
		product.ExpirationDate = parsedDate
		products = append(products, product)
	}
	return products, nil
}

// Create adds a new product to the repository.
func (p *ProductSQL) Create(product domain.Product) (domain.Product, error) {
	// insert product into db
	result, err := p.db.Exec(
		queries.CreateProduct,
		product.ProductCode, product.Description, product.Width, product.Height, product.Length,
		product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature,
		product.FreezingRate, product.ProductTypeId, product.SellerId, product.ExpirationDate,
	)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[string]error{
			"uq_product_code": utils.ErrProductCodeAlreadyExists,
			"product_type_id": utils.ErrProductTypeNotFound,
			"seller_id":       utils.ErrSellerNotFound,
		}

		for cnt, customErr := range dbErrorsMap {
			if utils.ValidateConstraintFailed(err, cnt) {
				return domain.Product{}, customErr
			}
		}

		return domain.Product{}, utils.ErrInternalServer
	}

	if result == nil {
		return domain.Product{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()

	if err != nil {
		return domain.Product{}, utils.ErrInternalServer
	}

	product.Id = id

	return product, nil

}

// GetById returns a product by its ID from the repository.
func (p *ProductSQL) GetById(id int64) (domain.Product, error) {
	var product domain.Product
	var expirationDate []uint8

	row := p.db.QueryRow(queries.GetProductById, id)
	err := row.Scan(
		&product.Id, &product.ProductCode, &product.Description, &product.Width, &product.Height,
		&product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature,
		&product.FreezingRate, &product.ProductTypeId, &product.SellerId, &expirationDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, utils.ErrProductNotFound
		}
		return domain.Product{}, utils.ErrInternalServer
	}

	parsedDate, err := utils.ParseDateTime(expirationDate)
	if err != nil {
		return domain.Product{}, utils.ErrFailedParseDate
	}
	product.ExpirationDate = parsedDate

	return product, nil
}

// Update updates a product in the repository.
func (p *ProductSQL) Update(product domain.Product) error {
	// update product in db
	_, err := p.db.Exec(
		queries.UpdateProduct,
		product.ProductCode, product.Description, product.Width, product.Height, product.Length,
		product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature,
		product.FreezingRate, product.ProductTypeId, product.SellerId, product.ExpirationDate,
		product.Id,
	)
	if err != nil {
		if dbErr := utils.ValidateErrorTypeSQL(err); dbErr != nil {
			dbErrorsMap := map[string]error{
				"uq_product_code": utils.ErrProductCodeAlreadyExists,
				"product_type_id": utils.ErrProductTypeNotFound,
				"seller_id":       utils.ErrSellerNotFound,
			}

			for cnt, customErr := range dbErrorsMap {
				if utils.ValidateConstraintFailed(err, cnt) {
					return customErr
				}
			}
		}
	}

	return nil
}

// Delete removes a product from the repository.
func (p *ProductSQL) Delete(id int64) error {
	result, err := p.db.Exec(queries.DeleteProduct, id)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBCannotDeleteOrUpdateParent) {
			return utils.ErrNoDeleteByForeignKey
		}
		return utils.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrInternalServer
	}

	if rowsAffected == 0 {
		return utils.ErrProductNotFound
	}

	return nil
}

func (r *ProductSQL) FindAllReports() (reportProducts []domain.ReportProductRecord, err error) {
	rows, err := r.db.Query(queries.FindAllReports)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)
	for rows.Next() {
		var reportProductRecord domain.ReportProductRecord
		err = rows.Scan(&reportProductRecord.ProductId, &reportProductRecord.Description, &reportProductRecord.RecordsCount)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		reportProducts = append(reportProducts, reportProductRecord)
	}
	return reportProducts, nil
}

func (r *ProductSQL) FindReportById(id int64) (reportProductRecord domain.ReportProductRecord, err error) {
	row := r.db.QueryRow(queries.FindReportByID, id)
	err = row.Scan(&reportProductRecord.ProductId, &reportProductRecord.Description, &reportProductRecord.RecordsCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ReportProductRecord{}, utils.ErrProductNotFound
		}
		return domain.ReportProductRecord{}, utils.ErrInternalServer
	}
	return reportProductRecord, nil
}
