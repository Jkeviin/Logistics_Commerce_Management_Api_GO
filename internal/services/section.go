package services

import (
	"errors"
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// SectionDefault is a struct that represents the default implementation sections
type SectionDefault struct {
	//rp is the repository that will be used by the service
	rp repository.SectionRepository
}

// NewSectionDefault is a function that returns a new instance of SectionDefault
func NewSectionDefault(rp repository.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

// FindAll is a method that returns a slice of all sections
func (sv *SectionDefault) FindAll() (sections []domain.Section, err error) {
	sections, err = sv.rp.FindAll()
	if err != nil {
		sections = nil
		err = utils.ErrInternalServer
	}
	return
}

// FindById is a method that returns the section that matches the given id
func (sv *SectionDefault) FindById(id int64) (section domain.Section, err error) {
	section, err = sv.rp.FindById(id)
	if err != nil {
		section = domain.Section{}
		err = utils.ErrSectionNotFound
	}
	return
}

// Create is a method that creates a new section
func (sv *SectionDefault) Create(sectionIn domain.Section) (sectionRes domain.Section, err error) {

	err = utils.ValidateStructGoValidator(sectionIn)
	if err != nil {
		sectionRes = domain.Section{}
		err = utils.ErrBadRequest
		return
	}

	sectionRes, err = sv.rp.Create(sectionIn)
	if err != nil {
		sectionRes = domain.Section{}
		if errors.Is(err, utils.ErrSectionAlreadyExists) {
			return
		}
		err = utils.ErrInternalServer
	}
	return
}

// Update is a method that updates a section in the repository
func (s *SectionDefault) Update(id int64, sectionPatch domain.SectionPatchAttributes) (section domain.Section, err error) {
	// Validate if the section exists and get the section to update
	section, err = s.rp.FindById(id)
	if err != nil {
		return domain.Section{}, utils.ErrSectionNotFound
	}
	// Update the section with the new data
	utils.UpdateStruct(&section, &sectionPatch)

	// Validate the section data with the govalidator
	if err := utils.ValidateStructGoValidator(section); err != nil {
		return domain.Section{}, fmt.Errorf("Error al validar la sección: %w", err)
	}

	// Save the updated section in the repository
	section, err = s.rp.Update(id, section)
	if err != nil {
		if errors.Is(err, utils.ErrSectionAlreadyExists) {
			return
		}
		err = utils.ErrInternalServer
		section = domain.Section{}
		return
	}
	return
}

// Delete is a method that eliminates a section from the repository
func (s *SectionDefault) Delete(id int64) (err error) {
	if _, err := s.rp.FindById(id); err != nil {
		return utils.ErrSectionNotFound
	}
	if err := s.rp.Delete(id); err != nil {
		return utils.ErrInternalServer
	}

	return
}
