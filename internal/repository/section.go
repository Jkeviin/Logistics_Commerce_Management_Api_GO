package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// SectionMap is a struct that represents a map of Sections
type SectionMap struct {
	db     map[int64]domain.Section
	loader *loader.JSONLoader[domain.SectionDoc, int64]
}

// NewSectionMap is a function that returns a new instance of SectionMap
func NewSectionMap(db map[int64]domain.Section, loader *loader.JSONLoader[domain.SectionDoc, int64]) *SectionMap {
	defaultDb := make(map[int64]domain.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionMap{
		db:     defaultDb,
		loader: loader,
	}
}

// FindAll is a method that returns a slice of all sections
func (r *SectionMap) FindAll() (sections []domain.Section, err error) {
	sections = utils.MapToSlice(r.db)
	err = nil
	return
}

// FindById is a method that returns the section that matches the given id
func (r *SectionMap) FindById(id int64) (section domain.Section, err error) {
	section, ok := r.db[id]
	if !ok {
		err = utils.ErrSectionNotFound
		return
	}
	return
}

// Create is a method that creates a new section
func (r *SectionMap) Create(sectionIn domain.Section) (sectionRes domain.Section, err error) {
	sectionIn.ID = time.Now().UnixNano()
	for _, section := range r.db {
		if section.SectionNumber == sectionIn.SectionNumber {
			err = utils.ErrSectionAlreadyExists
			return
		}
	}

	r.db[sectionIn.ID] = sectionIn

	if err := r.loader.SaveToJSON(r.getDBInSectionDoc()); err != nil {
		return domain.Section{}, err
	}

	sectionRes = sectionIn
	return
}

// Update is a method that updates a section in the repository
func (r *SectionMap) Update(id int64, sectionIn domain.Section) (sectionRes domain.Section, err error) {
	if _, exists := r.db[id]; !exists {
		return domain.Section{}, utils.ErrSectionNotFound
	}

	//validate if the fields sectionNumber already exists
	for _, section := range r.db {
		if section.SectionNumber == sectionIn.SectionNumber && section.ID != id {
			err = utils.ErrSectionAlreadyExists
			return
		}
	}

	sectionIn.ID = id
	r.db[id] = sectionIn

	sectionRes = sectionIn
	err = r.saveToJSON()
	return
}

// Delete is a method that eliminates a section from the repository
func (r *SectionMap) Delete(id int64) (err error) {
	if _, exists := r.db[id]; !exists {
		return utils.ErrSectionNotFound
	}

	delete(r.db, id)
	err = r.saveToJSON()
	return
}

// AUXILIAR FUNCTIONS
func (r *SectionMap) getDBInSectionDoc() (db map[int64]domain.SectionDoc) {
	db = make(map[int64]domain.SectionDoc)
	for key, value := range r.db {
		db[key] = value.ParseToSectionDoc()
	}
	return
}

func (r *SectionMap) saveToJSON() error {
	return r.loader.SaveToJSON(r.toSectionDoc())
}

func (r *SectionMap) toSectionDoc() map[int64]domain.SectionDoc {
	dbDoc := make(map[int64]domain.SectionDoc, len(r.db))
	for id, section := range r.db {
		dbDoc[id] = section.ParseToSectionDoc()
	}
	return dbDoc
}
