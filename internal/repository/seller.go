package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain/filters"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// NewSellerSQL creates a new instance of SellerSQL
func NewSellerSQL(db *sql.DB) repository.SellerRepository {
	return &SellerSQL{
		db: db,
	}
}

// SellerSQL represents the seller repository
type SellerSQL struct {
	db *sql.DB
}

// FindAll returns all sellers
func (repo *SellerSQL) FindAll() ([]domain.Seller, error) {
	rows, err := repo.db.Query(queries.FindSellerBase)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)
	return mapRowsToSellers(rows)
}

// FindByFilter returns a seller matching the filter
func (repo *SellerSQL) FindByFilter(filter filters.SellerFilter) (domain.Seller, error) {
	if !filter.HasFilters() {
		return domain.Seller{}, utils.ErrInvalidFilter
	}
	query := queries.FindSellerBase
	where, args := utils.BuildSQLWhereClause(&filter, filter.FieldColumnMap())
	row := repo.db.QueryRow(query+where, args...)
	return mapRowToSeller(row)
}

// Create inserts a new seller
func (repo *SellerSQL) Create(seller domain.Seller) (domain.Seller, error) {
	result, err := repo.db.Exec(
		queries.CreateSeller,
		seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityID,
	)
	if dbErr := utils.ValidateErrorTypeSQL(err); dbErr != nil {
		dbMap := []error{utils.ErrDBDuplicateEntry, utils.ErrDBForeignKey}
		domainMap := []error{utils.ErrSellerConflict, utils.ErrLocalityForeignKey}
		for i, e := range dbMap {
			if utils.ErrorInSlice([]error{e}, dbErr) {
				return domain.Seller{}, domainMap[i]
			}
		}
		return domain.Seller{}, utils.ErrInternalServer
	}
	if result == nil {
		return domain.Seller{}, utils.ErrInternalServer
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Seller{}, utils.ErrInternalServer
	}
	seller.ID = id
	return seller, nil
}

// Update updates an existing seller
func (repo *SellerSQL) Update(id int64, seller domain.Seller) (domain.Seller, error) {
	_, err := repo.db.Exec(
		queries.UpdateSeller,
		seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityID, id,
	)
	if dbErr := utils.ValidateErrorTypeSQL(err); dbErr != nil {
		dbMap := []error{utils.ErrDBDuplicateEntry, utils.ErrDBForeignKey}
		domainMap := []error{utils.ErrSellerConflict, utils.ErrLocalityForeignKey}
		for i, e := range dbMap {
			if utils.ErrorInSlice([]error{e}, dbErr) {
				return domain.Seller{}, domainMap[i]
			}
		}
		return domain.Seller{}, utils.ErrInternalServer
	}

	seller.ID = id
	return seller, nil
}

// Delete removes a seller
func (repo *SellerSQL) Delete(id int64) error {
	result, err := repo.db.Exec(queries.DeleteSeller, id)
	if err != nil {
		dbErr := utils.ValidateErrorTypeSQL(err)
		if utils.ErrorInSlice([]error{utils.ErrDBForeignKey}, dbErr) {
			return utils.ErrLocalityForeignKey
		}
		return utils.ErrInternalServer
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return utils.ErrInternalServer
	}
	if affectedRows == 0 {
		return utils.ErrSellerNotFound
	}
	return nil
}

// --- Private helper functions ---

func mapRowToSeller(row *sql.Row) (domain.Seller, error) {
	var sellerDomain domain.Seller
	err := row.Scan(&sellerDomain.ID, &sellerDomain.CID, &sellerDomain.CompanyName, &sellerDomain.Address, &sellerDomain.Telephone, &sellerDomain.LocalityID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Seller{}, utils.ErrSellerNotFound
	}
	if err != nil {
		return domain.Seller{}, utils.ErrInternalServer
	}
	return sellerDomain, nil
}

func mapRowsToSellers(rows *sql.Rows) ([]domain.Seller, error) {
	sellerList := make([]domain.Seller, 0)
	for rows.Next() {
		var sellerRecord domain.Seller
		err := rows.Scan(&sellerRecord.ID, &sellerRecord.CID, &sellerRecord.CompanyName, &sellerRecord.Address, &sellerRecord.Telephone, &sellerRecord.LocalityID)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		sellerList = append(sellerList, sellerRecord)
	}
	return sellerList, nil
}
