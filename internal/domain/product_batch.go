package domain

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type ProductBatch struct {
	ID                 int64       `valid:"-"`
	BatchNumber        *string     `valid:"required~El número de lote es obligatorio"`
	CurrentQuantity    *int        `valid:"notNil~La cantidad actual es obligatoria,range(0|2147483647)~La cantidad actual debe ser mayor o igual a 0"`
	CurrentTemperature *float64    `valid:"notNil~La temperatura actual es obligatoria,float19_2~La temperatura actual no es compatible con el formato de 19 enteros y 2 decimales"`
	DueDate            *utils.Date `valid:"notNil~La fecha de vencimiento es obligatoria"`
	InitialQuantity    *int        `valid:"notNil~La cantidad inicial es obligatoria,range(0|2147483647)~La cantidad inicial debe ser mayor o igual a 1"`
	ManufacturingDate  *utils.Date `valid:"notNil~La fecha de fabricación es obligatoria"`
	ManufacturingHour  *int        `valid:"notNil~La hora de fabricación es obligatoria"`
	MinimumTemperature *float64    `valid:"notNil~La temperatura mínima es obligatoria,float19_2~La temperatura mínima no es compatible con el formato de 19 enteros y 2 decimales"`
	ProductID          *int64      `valid:"required~El ID del producto es obligatorio"`
	SectionID          *int64      `valid:"required~El ID de la sección es obligatorio"`
}

type ProductBatchAttributes struct {
	BatchNumber        *string     `json:"batch_number"`
	CurrentQuantity    *int        `json:"current_quantity"`
	CurrentTemperature *float64    `json:"current_temperature"`
	DueDate            *utils.Date `json:"due_date"`
	InitialQuantity    *int        `json:"initial_quantity"`
	ManufacturingDate  *utils.Date `json:"manufacturing_date"`
	ManufacturingHour  *int        `json:"manufacturing_hour"`
	MinimumTemperature *float64    `json:"minimum_temperature"`
	ProductID          *int64      `json:"product_id"`
	SectionID          *int64      `json:"section_id"`
}

type ProductBatchDoc struct {
	ID int64 `json:"id"`
	ProductBatchAttributes
}

type ProductBatchResponse struct {
	Data ProductBatchDoc `json:"data"`
}

type ProductBatchesResponse struct {
	Data []ProductBatchDoc `json:"data"`
}

func (p *ProductBatch) ParseToProductBatchDoc() ProductBatchDoc {
	return ProductBatchDoc{
		ID: p.ID,
		ProductBatchAttributes: ProductBatchAttributes{
			BatchNumber:        p.BatchNumber,
			CurrentQuantity:    p.CurrentQuantity,
			CurrentTemperature: p.CurrentTemperature,
			DueDate:            p.DueDate,
			InitialQuantity:    p.InitialQuantity,
			ManufacturingDate:  p.ManufacturingDate,
			ManufacturingHour:  p.ManufacturingHour,
			MinimumTemperature: p.MinimumTemperature,
			ProductID:          p.ProductID,
			SectionID:          p.SectionID,
		},
	}
}

func (p *ProductBatchDoc) ParseToProductBatch() ProductBatch {
	productBatch := p.ProductBatchAttributes.ParseToProductBatch()
	productBatch.ID = p.ID
	return productBatch
}

func (p *ProductBatchAttributes) ParseToProductBatch() ProductBatch {
	return ProductBatch{
		BatchNumber:        p.BatchNumber,
		CurrentQuantity:    p.CurrentQuantity,
		CurrentTemperature: p.CurrentTemperature,
		DueDate:            p.DueDate,
		InitialQuantity:    p.InitialQuantity,
		ManufacturingDate:  p.ManufacturingDate,
		ManufacturingHour:  p.ManufacturingHour,
		MinimumTemperature: p.MinimumTemperature,
		ProductID:          p.ProductID,
		SectionID:          p.SectionID,
	}
}
