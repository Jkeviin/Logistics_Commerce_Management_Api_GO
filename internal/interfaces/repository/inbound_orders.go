package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type InboundOrdersRepository interface {
	Create(inboundOrder domain.InboundOrder) (domain.InboundOrder, error)
}
