package queries

const (
	GetAllProducts = `
		SELECT id, product_code, description, width, height, length, netweight, expiration_rate, 
			   recommended_freezing_temperature, freezing_rate, product_type_id, seller_id, expiration_date
		FROM product
	`

	CreateProduct = `
		INSERT INTO product (product_code, description, width, height, length, netweight, expiration_rate,
			recommended_freezing_temperature, freezing_rate, product_type_id, seller_id, expiration_date)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	GetProductById = `
		SELECT id, product_code, description, width, height, length, netweight, expiration_rate,
			   recommended_freezing_temperature, freezing_rate, product_type_id, seller_id, expiration_date
		FROM product
		WHERE id = ? 
	`
	UpdateProduct = `
		UPDATE product
		SET product_code = ?, description = ?, width = ?, height = ?, length = ?, netweight = ?,
			expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?,
			product_type_id = ?, seller_id = ?, expiration_date = ?
		WHERE id = ?
	`
	DeleteProduct = `
		DELETE FROM product
		WHERE id = ?
	`
	FindAllReports = `select p.id, p.description, count(pr.id) from product_record pr 
		join product p ON p.id = pr.product_id
		group by p.id;`
	FindReportByID = `select p.id, p.description, count(pr.id) from product_record pr 
		join product p ON p.id = pr.product_id
		where p.id = ? group by p.id;`
)
