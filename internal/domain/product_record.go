package domain

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

type ProductRecord struct {
	Id             int64
	ProductId      *int        `valid:"notNil~El id de producto es obligatorio,range(0|2147483647)~El id del producto debe ser mayor o igual a 0"`
	LastUpdateDate *utils.Date `valid:"notNil~La fecha de actualización es obligatoria"`
	PurchasePrice  *float64    `valid:"notNil~El precio de compra es obligatorio,float19_2~El precio de compra debe tener hasta 19 dígitos y 2 decimales"`
	SalePrice      *float64    `valid:"notNil~El precio de venta es obligatorio,float19_2~El precio de venta debe tener hasta 19 dígitos y 2 decimales"`
}

type ProductRecordDoc struct {
	Id int64 `json:"id"`
	ProductRecordAttributes
}

type ProductRecordAttributes struct {
	ProductId      *int        `json:"product_id"`
	LastUpdateDate *utils.Date `json:"last_update_date"`
	PurchasePrice  *float64    `json:"purchase_price"`
	SalePrice      *float64    `json:"sale_price"`
}

type ProductRecordResponse struct {
	Data ProductRecordDoc `json:"data"`
}

func (p *ProductRecordAttributes) ParseToProductRecord() ProductRecord {
	return ProductRecord{
		ProductId:      p.ProductId,
		LastUpdateDate: p.LastUpdateDate,
		PurchasePrice:  p.PurchasePrice,
		SalePrice:      p.SalePrice,
	}
}

func (p *ProductRecord) ParseToProductRecordDoc() ProductRecordDoc {
	return ProductRecordDoc{
		Id: p.Id,
		ProductRecordAttributes: ProductRecordAttributes{
			ProductId:      p.ProductId,
			LastUpdateDate: p.LastUpdateDate,
			PurchasePrice:  p.PurchasePrice,
			SalePrice:      p.SalePrice,
		},
	}
}
