package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type PurchaseOrderRepository interface {
	// Create inserts a new purchase order into the repository
	Create(po domain.PurchaseOrder) (domain.PurchaseOrder, error)
}
