package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewEmployeeSQL(db *sql.DB) *EmployeeSQL {
	return &EmployeeSQL{
		db: db,
	}
}

type EmployeeSQL struct {
	db *sql.DB
}

func (r *EmployeeSQL) FindAll() ([]domain.Employee, error) {
	rows, err := r.db.Query(queries.GetAllEmployees)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	var employees []domain.Employee

	for rows.Next() {
		var employee domain.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.CardNumberId,
			&employee.FirstName,
			&employee.LastName,
			&employee.WarehouseId,
		)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *EmployeeSQL) Find(id int64) (domain.Employee, error) {
	row := r.db.QueryRow(queries.GetEmployeeById, id)
	var employee domain.Employee
	err := row.Scan(
		&employee.Id,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Employee{}, utils.ErrEmployeeNotFound
		}
		return domain.Employee{}, utils.ErrInternalServer
	}
	return employee, nil
}

func (r *EmployeeSQL) Create(employee domain.Employee) (domain.Employee, error) {
	result, err := r.db.Exec(
		queries.CreateEmployee,
		employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrCardNumberAlreadyExists,
			utils.ErrDBForeignKey:     utils.ErrWareHouseNotFound,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Employee{}, dbErr
		}
		return domain.Employee{}, utils.ErrInternalServer
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Employee{}, utils.ErrInternalServer
	}
	employee.Id = id
	return employee, nil

}

func (r *EmployeeSQL) Delete(id int64) error {
	result, err := r.db.Exec(queries.DeleteEmployee, id)
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
		return utils.ErrEmployeeNotFound
	}

	return nil

}

func (r *EmployeeSQL) Update(id int64, employee domain.Employee) (domain.Employee, error) {
	_, err := r.db.Exec(
		queries.UpdateEmployee,
		employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId, id,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {

		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrCardNumberAlreadyExists,
			utils.ErrDBForeignKey:     utils.ErrWareHouseNotFound,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Employee{}, dbErr
		}

		return domain.Employee{}, utils.ErrInternalServer
	}

	return employee, nil
}

func (r *EmployeeSQL) ReportInboundOrders() ([]domain.EmployeeWithInboundOrdersCount, error) {
	rows, err := r.db.Query(queries.GetEmployeesWithInboundOrdersCount)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	employees := make([]domain.EmployeeWithInboundOrdersCount, 0)

	for rows.Next() {
		var emplWithCount domain.EmployeeWithInboundOrdersCount
		err = rows.Scan(
			&emplWithCount.ID,
			&emplWithCount.CardNumber,
			&emplWithCount.FirstName,
			&emplWithCount.LastName,
			&emplWithCount.WarehouseId,
			&emplWithCount.InboundOrdersCount,
		)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		employees = append(employees, emplWithCount)
	}
	return employees, nil
}

func (r *EmployeeSQL) ReportInboundOrdersById(id int64) (domain.EmployeeWithInboundOrdersCount, error) {
	row := r.db.QueryRow(queries.GetEmployeeWithInboundOrdersCountById, id)
	var employee domain.EmployeeWithInboundOrdersCount
	err := row.Scan(
		&employee.ID,
		&employee.CardNumber,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
		&employee.InboundOrdersCount,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.EmployeeWithInboundOrdersCount{}, utils.ErrEmployeeNotFound
		}
		return domain.EmployeeWithInboundOrdersCount{}, utils.ErrInternalServer
	}
	return employee, nil
}
