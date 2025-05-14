package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type LocalityService interface {
	FindSellers(id string) (domain.LocalityReportAttributes, error)
	Create(locality domain.Locality) (domain.Locality, error)
	GetCantCarriesPerLocality(id string) ([]domain.LocalityWithCount, error)
	FindAllSellers() ([]domain.LocalityReportAttributes, error)
}
