package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewLocalitySQL(db *sql.DB) *LocalitySQL {
	return &LocalitySQL{
		db: db,
	}
}

type LocalitySQL struct {
	db *sql.DB
}

func (r *LocalitySQL) FindSellers(id string) (domain.LocalityReportAttributes, error) {
	row := r.db.QueryRow(queries.GetCantCarriesPerLocality, id)
	var locality domain.LocalityReportAttributes
	err := row.Scan(
		&locality.LocalityID,
		&locality.LocalityName,
		&locality.SellersCount,
	)
	fmt.Println(err)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return domain.LocalityReportAttributes{}, utils.ErrSellersNotFound
		}
		return domain.LocalityReportAttributes{}, utils.ErrInternalServer
	}
	return locality, nil
}

func (r *LocalitySQL) FindAllSellers() ([]domain.LocalityReportAttributes, error) {
	rows, err := r.db.Query(queries.GetAllCarriesPerLocality2)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	var carries []domain.LocalityReportAttributes

	for rows.Next() {
		var carrie domain.LocalityReportAttributes
		err := rows.Scan(
			&carrie.LocalityID,
			&carrie.LocalityName,
			&carrie.SellersCount,
		)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		carries = append(carries, carrie)
	}
	return carries, nil
}

func (r *LocalitySQL) Create(locality domain.Locality) (domain.Locality, error) {
	_, err := r.db.Exec(
		queries.Createlocality,
		locality.Id, locality.LocalityName, locality.ProvinceName, locality.CountryName,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[error]error{
			utils.ErrDBDuplicateEntry: utils.ErrIdAlreadyExist,
		}
		if dbErr, ok := dbErrorsMap[errDB]; ok {
			return domain.Locality{}, dbErr
		}
		return domain.Locality{}, utils.ErrInternalServer
	}

	return locality, nil

}

// GetAllcarriesPerLocality retrieves all carries per locality.
func (r *LocalitySQL) GetAllcarriesPerLocality() ([]domain.LocalityWithCount, error) {
	// Ejecuta la consulta para obtener los datos
	rows, err := r.db.Query(queries.GetAllCarriesPerLocality)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows) // Asegura que las filas se cierren después de procesarlas

	var localities []domain.LocalityWithCount

	// Itera sobre el conjunto de resultados
	for rows.Next() {
		var locality domain.LocalityWithCount

		// Escanea la fila actual dentro de la estructura locality
		err := rows.Scan(
			&locality.Id,
			&locality.LocalityName,
			&locality.CantCarries,
		)
		if err != nil {
			return nil, utils.ErrInternalServer
		}

		// Agrega la localidad al slice de resultados
		localities = append(localities, locality)
	}

	return localities, nil // Retorna la lista de localidades con el conteo de carries
}

// GetCantCarriesPerLocality retrieves the count of carries per locality by locality ID.
func (r *LocalitySQL) GetCantCarriesPerLocality(id string) (domain.LocalityWithCount, error) {
	row := r.db.QueryRow(queries.GetCantCarriesPerLocality, id) // Usando una consulta para obtener un solo resultado
	var localityWithCount domain.LocalityWithCount
	err := row.Scan(
		&localityWithCount.Id,
		&localityWithCount.LocalityName,
		&localityWithCount.CantCarries,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.LocalityWithCount{}, utils.ErrLocalityForeignKey // Error más específico
		}
		return domain.LocalityWithCount{}, utils.ErrInternalServer
	}
	return localityWithCount, nil
}
