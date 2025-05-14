package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type BuyerRepository interface {
	FindAll() ([]domain.Buyer, error)
	Create(buyer domain.Buyer) (domain.Buyer, error)
	FindById(id int64) (buyer domain.Buyer, err error)
	Update(id int64, buyer domain.Buyer) (domain.Buyer, error)
	Delete(id int64) (err error)
	ReportPurchaseOrdersByBuyerID(buyerID int64) ([]domain.BuyerPurchaseOrdersReport, error)
	ReportPurchaseOrdersAll() ([]domain.BuyerPurchaseOrdersReport, error)
}
