package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type WarehouseSQL struct {
	db *sql.DB
}

func NewWarehouseSQL(db *sql.DB) *WarehouseSQL {
	return &WarehouseSQL{db: db}
}

func (r *WarehouseSQL) FindAll() ([]domain.Warehouse, error) {
	rows, err := r.db.Query(queries.FindAllWarehouses)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	var warehouses []domain.Warehouse
	for rows.Next() {
		var warehouse domain.Warehouse
		if err := rows.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimumCapacity, &warehouse.MinimumTemperature, &warehouse.LocalityID); err != nil {
			return nil, utils.ErrInternalServer
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (r *WarehouseSQL) Find(id int64) (domain.Warehouse, error) {
	row := r.db.QueryRow(queries.FindWarehouseByID, id)

	var warehouse domain.Warehouse
	if err := row.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimumCapacity, &warehouse.MinimumTemperature, &warehouse.LocalityID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Warehouse{}, utils.ErrWareHouseNotFound
		}
		return domain.Warehouse{}, utils.ErrInternalServer
	}

	return warehouse, nil
}

func (r *WarehouseSQL) Create(warehouse domain.Warehouse) (domain.Warehouse, error) {
	result, err := r.db.Exec(
		queries.CreateWarehouse,
		warehouse.WarehouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimumCapacity, warehouse.MinimumTemperature, warehouse.LocalityID,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrWarehouseCodeAlreadyExists,
			utils.ErrDBForeignKey:     utils.ErrLocalityForeignKey,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Warehouse{}, dbErr
		}
		return domain.Warehouse{}, utils.ErrInternalServer
	}

	if result == nil {
		return domain.Warehouse{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Warehouse{}, utils.ErrInternalServer
	}
	warehouse.ID = id
	return warehouse, nil
}

func (r *WarehouseSQL) Delete(id int64) error {
	result, err := r.db.Exec(queries.DeleteWarehouse, id)
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
		return utils.ErrWareHouseNotFound
	}

	return nil
}

func (r *WarehouseSQL) Update(id int64, warehouse domain.Warehouse) (domain.Warehouse, error) {
	_, err := r.db.Exec(
		queries.UpdateWarehouse,
		warehouse.WarehouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimumCapacity, warehouse.MinimumTemperature, warehouse.LocalityID, id,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrWarehouseCodeAlreadyExists,
			utils.ErrDBForeignKey:     utils.ErrLocalityForeignKey,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Warehouse{}, dbErr
		}
		return domain.Warehouse{}, utils.ErrInternalServer
	}

	return warehouse, nil
}
