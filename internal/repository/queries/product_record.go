package queries

const (
	CreateProductRecord = `
		insert into product_record (last_update_date, purchase_price, sale_price, product_id) 
		values (?, ?, ?, ?);`
)
