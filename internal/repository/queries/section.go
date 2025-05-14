package queries

const (
	FindAllSections = `
		SELECT id, 
			section_number, 
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id 
		FROM section
	`

	FindSectionById = `
		SELECT id,
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id 
		FROM section s 
		WHERE s.id = ?
	`

	CreateSection = `
		INSERT INTO section(
			section_number,
    		current_temperature,
    		minimum_temperature,
    		current_capacity,
    		minimum_capacity,
   			maximum_capacity,
    		warehouse_id,
    		product_type_id)
		VALUES 
			(?,?,?,?,?,?,?,?)
	`

	UpdateSection = `
		UPDATE section
		SET
			section_number = ?,
    		current_temperature = ?,
			minimum_temperature = ?,
			current_capacity = ?,
    		minimum_capacity = ?,
			maximum_capacity = ?,
			warehouse_id = ?,
			product_type_id = ?
		WHERE id = ?
	`

	DeleteSection = `
		DELETE FROM section
		WHERE id = ?
	`

	FindSectionWithProductCountById = `
		SELECT s.id, s.section_number, COUNT(p.id)
		FROM section s
		JOIN product_batch pb on pb.section_id = s.id
		JOIN product p on p.id = pb.product_id
		WHERE s.id = ?
		GROUP BY s.id
	`

	FindSectionWithProductCount = `
		SELECT s.id, s.section_number, COUNT(p.id)
		FROM section s
		JOIN product_batch pb on pb.section_id = s.id
		JOIN product p on p.id = pb.product_id
		GROUP BY s.id
	`
)
