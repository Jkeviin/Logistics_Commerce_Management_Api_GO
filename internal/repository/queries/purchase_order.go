package queries

const (
	CreatePurchaseOrder = `
		INSERT INTO purchase_order (order_number, order_date, tracking_code, buyer_id, product_record_id)
		VALUES (?, ?, ?, ?, ?)
	`
)
