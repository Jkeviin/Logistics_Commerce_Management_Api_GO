package queries

const (
	CreateCarrie = `
		INSERT INTO carry (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)
	`
)
