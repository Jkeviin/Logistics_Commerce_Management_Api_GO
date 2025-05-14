package queries

const (
	FindSellerBase = `
		SELECT id, cid, company_name, address, telephone, locality_id
		FROM seller
	`

	CreateSeller = `
		INSERT INTO seller (cid, company_name, address, telephone, locality_id)
		VALUES (?, ?, ?, ?, ?)
	`

	UpdateSeller = `
		UPDATE seller
		SET cid = ?, company_name = ?, address = ?, telephone = ?, locality_id = ?
		WHERE id = ?
	`

	DeleteSeller = `
		DELETE FROM seller
		WHERE id = ?
	`
)
