package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewLocalityDefault(rp repository.LocalityRepository) *LocalityDefault {
	return &LocalityDefault{rp: rp}
}

type LocalityDefault struct {
	rp repository.LocalityRepository
}

func (s *LocalityDefault) FindSellers(id string) (domain.LocalityReportAttributes, error) {
	return s.rp.FindSellers(id)
}

func (s *LocalityDefault) FindAllSellers() ([]domain.LocalityReportAttributes, error) {
	return s.rp.FindAllSellers()
}

func (s *LocalityDefault) Create(locality domain.Locality) (domain.Locality, error) {
	err := utils.ValidateStructGoValidator(locality)
	if err != nil {
		return domain.Locality{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return s.rp.Create(locality)
}

// GetCantCarriesPerLocality obtiene la cantidad de carries por localidad según su ID.
func (s *LocalityDefault) GetCantCarriesPerLocality(id string) ([]domain.LocalityWithCount, error) {
	// Validar la entrada
	if id == "" {
		return s.rp.GetAllcarriesPerLocality()
	}

	// Llamar al método del repositorio
	localities, err := s.rp.GetCantCarriesPerLocality(id)
	if err != nil {
		return nil, err
	}
	return []domain.LocalityWithCount{localities}, nil
}
