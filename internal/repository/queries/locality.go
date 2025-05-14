package queries

const (
	Createlocality = `INSERT INTO locality (id, locality_name, province_name, country_name ) VALUES (?, ?, ?, ?)`

	GetCantCarriesPerLocality = `SELECT 
    l.id AS locality_id,
    l.locality_name,
    COUNT(s.id) AS sellers_count
	FROM locality l
	LEFT JOIN seller s ON l.id = s.locality_id
	WHERE l.id = ?
	GROUP BY l.id, l.locality_name;`

	GetAllCarriesPerLocality2 = `SELECT
    l.id AS locality_id,
    l.locality_name,
    COUNT(s.id) AS sellers_count
	FROM
		locality l
	LEFT JOIN
		seller s ON s.locality_id = l.id
	GROUP BY
		l.id, l.locality_name;`

	GetAllCarriesPerLocality = `SELECT 
		l.id AS locality_id, 
		l.locality_name,
		COUNT(c.id) AS carries_count
	FROM 
		locality l
	LEFT JOIN 
		carry c 
	ON 
		l.id = c.locality_id
	GROUP BY 
		l.id, l.locality_name, l.province_name, l.country_name;`
)
