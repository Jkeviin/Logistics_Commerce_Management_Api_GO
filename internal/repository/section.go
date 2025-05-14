package repository

import (
	"database/sql"
	"errors"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository/queries"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// SectionSQL is a struct that represents a map of Sections
type SectionSQL struct {
	db *sql.DB
}

// NewSectionSQL is a function that returns a new instance of SectionSQL
func NewSectionSQL(db *sql.DB) *SectionSQL {
	return &SectionSQL{
		db: db,
	}
}

// FindAll is a method that returns a slice of all sections
func (r *SectionSQL) FindAll() (sections []domain.Section, err error) {
	rows, err := r.db.Query(queries.FindAllSections)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	for rows.Next() {
		var section domain.Section
		err = rows.Scan(&section.ID,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseID,
			&section.ProductTypeID)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		sections = append(sections, section)
	}
	return
}

// FindById is a method that returns the section that matches the given id
func (r *SectionSQL) FindById(id int64) (section domain.Section, err error) {
	row := r.db.QueryRow(queries.FindSectionById, id)
	//Almaceno los datos de la row en la estructura
	err = row.Scan(&section.ID,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseID,
		&section.ProductTypeID)
	//Validaciones
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = utils.ErrSectionNotFound
			section = domain.Section{}
			return
		}
		section = domain.Section{}
		err = utils.ErrInternalServer
	}
	return
}

// Create is a method that creates a new section
func (r *SectionSQL) Create(sectionIn domain.Section) (sectionRes domain.Section, err error) {
	res, err := r.db.Exec(
		queries.CreateSection,
		sectionIn.SectionNumber, sectionIn.CurrentTemperature, sectionIn.MinimumTemperature, sectionIn.CurrentCapacity, sectionIn.MinimumCapacity, sectionIn.MaximumCapacity, sectionIn.WarehouseID, sectionIn.ProductTypeID,
	)

	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[string]error{
			"uq_section_number": utils.ErrSectionAlreadyExists,
			"product_type_id":   utils.ErrProductNotFound,
			"warehouse_id":      utils.ErrWareHouseNotFound,
		}

		for cnt, customErr := range dbErrorsMap {
			if utils.ValidateConstraintFailed(err, cnt) {
				return domain.Section{}, customErr
			}
		}

		return domain.Section{}, utils.ErrInternalServer
	}

	id, err := res.LastInsertId()
	if err != nil {
		err = utils.ErrInternalServer
		return
	}

	sectionRes = sectionIn
	sectionRes.ID = id
	err = nil
	return
}

// Update is a method that updates a section in the repository
func (r *SectionSQL) Update(id int64, sectionIn domain.Section) (sectionRes domain.Section, err error) {
	_, err = r.db.Exec(
		queries.UpdateSection,
		sectionIn.SectionNumber, sectionIn.CurrentTemperature, sectionIn.MinimumTemperature, sectionIn.CurrentCapacity, sectionIn.MinimumCapacity, sectionIn.MaximumCapacity, sectionIn.WarehouseID, sectionIn.ProductTypeID, id,
	)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		dbErrorsMap := map[string]error{
			"uq_section_number": utils.ErrSectionAlreadyExists,
			"product_type_id":   utils.ErrProductNotFound,
			"warehouse_id":      utils.ErrWareHouseNotFound,
		}
		for cnt, customErr := range dbErrorsMap {
			if utils.ValidateConstraintFailed(err, cnt) {
				return domain.Section{}, customErr
			}
		}
		return domain.Section{}, utils.ErrInternalServer
	}

	sectionRes = sectionIn
	sectionRes.ID = id
	err = nil
	return
}

// Delete is a method that eliminates a section from the repository
func (r *SectionSQL) Delete(id int64) (err error) {
	res, err := r.db.Exec(queries.DeleteSection, id)
	if errDB := utils.ValidateErrorTypeSQL(err); errDB != nil {
		if errors.Is(errDB, utils.ErrDBCannotDeleteOrUpdateParent) {
			return utils.ErrNoDeleteByForeignKey
		}
		return utils.ErrInternalServer
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.ErrInternalServer
	}
	if rowsAffected == 0 {
		return utils.ErrSectionNotFound
	}
	return
}

func (r *SectionSQL) ReportProductBySectionById(idSection int64) (domain.SectionWithProductCount, error) {
	row := r.db.QueryRow(queries.FindSectionWithProductCountById, idSection)
	var sectionWithProductCount domain.SectionWithProductCount
	err := row.Scan(
		&sectionWithProductCount.SectionID,
		&sectionWithProductCount.SectionNumber,
		&sectionWithProductCount.ProductCount,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.SectionWithProductCount{}, utils.ErrSectionNotFound
		}
		return domain.SectionWithProductCount{}, utils.ErrInternalServer
	}
	return sectionWithProductCount, nil
}

func (r *SectionSQL) ReportProductBySection() ([]domain.SectionWithProductCount, error) {
	rows, err := r.db.Query(queries.FindSectionWithProductCount)
	if err != nil {
		return nil, utils.ErrInternalServer
	}
	defer utils.CloseRows(rows)

	sectionsWithProducts := make([]domain.SectionWithProductCount, 0)

	for rows.Next() {
		var sectionWithProductCount domain.SectionWithProductCount
		err = rows.Scan(
			&sectionWithProductCount.SectionID,
			&sectionWithProductCount.SectionNumber,
			&sectionWithProductCount.ProductCount,
		)
		if err != nil {
			return nil, utils.ErrInternalServer
		}
		sectionsWithProducts = append(sectionsWithProducts, sectionWithProductCount)
	}
	return sectionsWithProducts, nil
}
