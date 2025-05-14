package domain

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"

type InboundOrder struct {
	ID             int64       `valid:"-"`
	OrderDate      *utils.Date `valid:"notNil~La fecha de la orden es obligatoria"`
	OrderNumber    *string     `valid:"required~El número de orden es obligatorio"`
	Temperature    *float64    `valid:"notNil~La temperatura es obligatoria,float19_2~La temperatura no es compatible con el formato de 19 enteros y 2 decimales"`
	EmployeeId     *int        `valid:"required~El ID del empleado es obligatorio"`
	ProductBatchId *int        `valid:"required~El ID del lote de productos es obligatorio"`
	WarehouseId    *int        `valid:"required~El ID del almacén es obligatorio"`
}

type InboundOrderAttributes struct {
	OrderDate      *utils.Date `json:"order_date"`
	OrderNumber    *string     `json:"order_number"`
	Temperature    *float64    `json:"temperature"`
	EmployeeId     *int        `json:"employee_id"`
	ProductBatchId *int        `json:"product_batch_id"`
	WarehouseId    *int        `json:"warehouse_id"`
}

type InboundOrderDoc struct {
	ID int64 `json:"id"`
	InboundOrderAttributes
}

type InboundOrderResponse struct {
	Data InboundOrderDoc `json:"data"`
}

type InboundOrdersResponse struct {
	Data []InboundOrderDoc `json:"data"`
}

func (i *InboundOrder) ParseToInboundOrderDoc() InboundOrderDoc {
	return InboundOrderDoc{
		ID: i.ID,
		InboundOrderAttributes: InboundOrderAttributes{
			OrderDate:      i.OrderDate,
			OrderNumber:    i.OrderNumber,
			Temperature:    i.Temperature,
			EmployeeId:     i.EmployeeId,
			ProductBatchId: i.ProductBatchId,
			WarehouseId:    i.WarehouseId,
		},
	}
}

func (i *InboundOrderAttributes) ParseToInboundOrder() InboundOrder {
	return InboundOrder{
		OrderDate:      i.OrderDate,
		OrderNumber:    i.OrderNumber,
		Temperature:    i.Temperature,
		EmployeeId:     i.EmployeeId,
		ProductBatchId: i.ProductBatchId,
		WarehouseId:    i.WarehouseId,
	}
}

func (i *InboundOrderDoc) ParseToInboundOrder() InboundOrder {
	inboundOrder := i.InboundOrderAttributes.ParseToInboundOrder()
	inboundOrder.ID = i.ID
	return inboundOrder
}
