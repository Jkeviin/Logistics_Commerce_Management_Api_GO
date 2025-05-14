package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type BuyerService interface {
	FindAll() ([]domain.Buyer, error)
	Create(buyer domain.Buyer) (domain.Buyer, error)
	FindById(id int64) (domain.Buyer, error)
	Update(id int64, buyer domain.BuyerAttributes) (domain.Buyer, error)
	Delete(id int64) (err error)
}
