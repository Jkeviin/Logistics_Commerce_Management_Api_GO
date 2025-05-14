package queries

const (
	CreateInboundOrder = `
		INSERT INTO inbound_order 
			(order_date, 
			order_number, 
			temperature, 
			employee_id, 
			product_batch_id, 
			warehouse_id) 
		VALUES 
			(?, ?, ?, ?, ?, ?)
	`
)
