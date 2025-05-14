package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type CarriesService interface {
	Create(product domain.Carries) (domain.Carries, error)
}
