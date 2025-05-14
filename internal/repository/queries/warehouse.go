package queries

const (
	FindAllWarehouses = `
		SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id 
		FROM warehouse
	`

	FindWarehouseByID = `
		SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id 
		FROM warehouse 
		WHERE id = ?
	`

	CreateWarehouse = `
		INSERT INTO warehouse (warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) 
		VALUES (?, ?, ?, ?, ?, ?)
	`

	DeleteWarehouse = `
		DELETE FROM warehouse
		WHERE id = ?
	`

	UpdateWarehouse = `
		UPDATE warehouse
		SET warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ?
		WHERE id = ?
	`
)
