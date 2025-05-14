package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

// SectionService is an interface that represents a section service
type SectionService interface {
	// FindAll returns a slice of all sections
	FindAll() (sections []domain.Section, err error)

	//FindById is a method that returns the section that matches the given id
	FindById(id int64) (section domain.Section, err error)

	//Create is a method that creates a new section
	Create(sectionIn domain.Section) (sectionRes domain.Section, err error)

	//Update is a method that updates a section in the repository
	Update(id int64, sectionPatch domain.SectionPatchAttributes) (sectionRes domain.Section, err error)

	// Delete is a method that eliminates a section from the repository
	Delete(id int64) (err error)
}
