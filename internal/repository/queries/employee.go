package queries

const (
	CreateEmployee = `INSERT INTO employee (id_card_number, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)`

	GetAllEmployees = `SELECT id, id_card_number, first_name, last_name, warehouse_id from employee`

	GetEmployeeById = `SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employee WHERE id = ?`

	DeleteEmployee = `
		DELETE FROM employee
		WHERE id = ?
		`
	UpdateEmployee = `
		UPDATE employee
		SET id_card_number = ?, first_name = ?, last_name = ?, warehouse_id = ?
		WHERE id = ?
	`

	GetEmployeeWithInboundOrdersCountById = `
		SELECT 
			e.id,
			e.id_card_number,	
    		e.first_name,
			e.last_name,
			e.warehouse_id,
    		COUNT(o.id) AS Total
		FROM employee e
		INNER JOIN inbound_order o ON o.employee_id = e.id
		WHERE e.id = ?
		GROUP BY e.id
	`

	GetEmployeesWithInboundOrdersCount = `
		SELECT
			e.id,
			e.id_card_number,
			e.first_name,
			e.last_name,
			e.warehouse_id,
			COUNT(o.id) AS Total
		FROM employee e
		INNER JOIN inbound_order o ON o.employee_id = e.id
		GROUP BY e.id
	`
)
