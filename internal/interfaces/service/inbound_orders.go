package service

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type InboundOrdersService interface {
	Create(inboundOrder domain.InboundOrder) (domain.InboundOrder, error)
}
