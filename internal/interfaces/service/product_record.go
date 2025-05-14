package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type ProductRecordService interface {
	Create(productRecord domain.ProductRecord) (domain.ProductRecord, error)
}
