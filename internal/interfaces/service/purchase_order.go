package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type PurchaseOrderService interface {
	Create(po domain.PurchaseOrder) (domain.PurchaseOrder, error)
}
