package queries

const (
	FindAllBuyers = `
		SELECT id, first_name, last_name, id_card_number
		FROM buyer
	`
	FindBuyerByID = `
		SELECT id, first_name, last_name, id_card_number
		FROM buyer
		WHERE id = ?
	`
	CreateBuyer = `INSERT INTO buyer (first_name, last_name, id_card_number) VALUES (?, ?, ?)`

	DeleteBuyer = `
		DELETE FROM buyer
		WHERE id = ?
	`

	UpdateBuyer = `
		UPDATE buyer
		SET first_name = ?, last_name = ?, id_card_number = ?
		WHERE id = ?
	`

	ReportPurchaseOrdersByBuyer = `
		SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(po.id) as purchase_orders
		FROM buyer b
		LEFT JOIN purchase_order po ON b.id = po.buyer_id
		GROUP BY b.id
	`

	ReportPurchaseOrdersByBuyerID = `
		SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(po.id) as purchase_orders
		FROM buyer b
		LEFT JOIN purchase_order po ON b.id = po.buyer_id
		WHERE b.id = ?
		GROUP BY b.id
	`
)
