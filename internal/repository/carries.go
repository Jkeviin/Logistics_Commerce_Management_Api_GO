package repository

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// ProductRepository is an interface that defines the methods for product repository.
func NewCarriesSQL(db *sql.DB) *CarriesSQL {
	return &CarriesSQL{
		db: db,
	}
}

// ProductSQL is a struct that implements the ProductRepository interface.
type CarriesSQL struct {
	db *sql.DB
}

// Create adds a new product to the repository.
func (c *CarriesSQL) Create(carries domain.Carries) (domain.Carries, error) {
	// insert carries into db
	result, err := c.db.Exec(queries.CreateCarrie, carries.CID, carries.CompanyName, carries.Address, carries.Telephone, carries.LocalityID)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrCarriesAlreadyExists,
			utils.ErrDBForeignKey:     utils.ErrLocalityForeignKey,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Carries{}, dbErr
		}
		return domain.Carries{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Carries{}, utils.ErrInternalServer
	}

	carries.ID = id
	return carries, nil
}
