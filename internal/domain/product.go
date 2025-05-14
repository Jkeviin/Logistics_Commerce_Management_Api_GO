package domain

import (
	"time"
)

// ProductAttributes encapsulates the common attributes of a product.

type ProductAttributes struct {
	ProductCode                    *string    `json:"product_code" valid:"required~El código del producto es obligatorio"`
	Description                    *string    `json:"description" valid:"required~La descripción es obligatoria"`
	Width                          *float64   `json:"width" valid:"required~El ancho es obligatorio"`
	Height                         *float64   `json:"height" valid:"required~La altura es obligatoria"`
	Length                         *float64   `json:"length" valid:"required~La longitud es obligatoria"`
	NetWeight                      *float64   `json:"netweight" valid:"required~El peso neto es obligatorio"`
	ExpirationRate                 *float64   `json:"expiration_rate" valid:"required~La tasa de caducidad es obligatoria"`
	RecommendedFreezingTemperature *float64   `json:"recommended_freezing_temperature" valid:"required~La temperatura de congelación recomendada es obligatoria"`
	FreezingRate                   *float64   `json:"freezing_rate" valid:"required~La tasa de congelación es obligatoria"`
	ProductTypeId                  *int       `json:"product_type_id" valid:"required~El ID del tipo de producto es obligatorio"`
	SellerId                       *int       `json:"seller_id" valid:"required~El ID del vendedor es obligatorio"`
	ExpirationDate                 *time.Time `json:"expiration_date,omitempty"`
}

// Product represents a product in the system.
type Product struct {
	Id int64 `json:"id" valid:"-"` // ID is not required for creation
	ProductAttributes
}

// ProductDoc represents a product document in the system.
type ProductDoc struct {
	ID int64 `json:"id"`
	ProductAttributes
}

// ProductResponse represents the response structure for a single product.
type ProductsResponse struct {
	Data []ProductDoc `json:"data"`
}

type ProductResponse struct {
	Data ProductDoc `json:"data"`
}

type ReportProductRecord struct {
	ProductId    int64  `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}

type ReportResponse struct {
	Data []ReportProductRecord `json:"data"`
}

// ParseToProductDoc converts a Product to a ProductDoc.
func (p *Product) ParseToProductDoc() ProductDoc {
	return ProductDoc{
		ID:                p.Id,
		ProductAttributes: p.ProductAttributes,
	}
}

// ParseToProduct converts a ProductDoc to a Product.
func (p *ProductDoc) ParseToProduct() Product {
	return Product{
		Id:                p.ID,
		ProductAttributes: p.ProductAttributes,
	}
}
