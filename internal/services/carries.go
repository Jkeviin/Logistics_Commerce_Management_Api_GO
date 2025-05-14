package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewCarriesDefault(rp repository.CarriesRepository) *CarriesDefault {
	return &CarriesDefault{rp: rp}
}

type CarriesDefault struct {
	rp repository.CarriesRepository
}

func (p *CarriesDefault) Create(carries domain.Carries) (domain.Carries, error) {
	//check all fields are not empty or invalid
	if err := utils.ValidateStructGoValidator(carries); err != nil {
		return domain.Carries{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	// Add the new product to the repository
	carrieCreated, err := p.rp.Create(carries)
	if err != nil {
		return domain.Carries{}, err
	}

	return carrieCreated, nil
}
