package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type ProductBatchService interface {
	Create(productBatch domain.ProductBatch) (domain.ProductBatch, error)
}
