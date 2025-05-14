package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type CarriesRepository interface {
	Create(product domain.Carries) (domain.Carries, error)
}
