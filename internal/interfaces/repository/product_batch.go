package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type ProductBatchRepository interface {
	Create(productBatch domain.ProductBatch) (domain.ProductBatch, error)
}
