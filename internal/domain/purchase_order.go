package domain

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// PurchaseOrder represents a purchase order in the domain
//go:generate mockgen -destination=../../mocks/domain_mock_purchaseorder.go -package=mocks . PurchaseOrderRepository

type PurchaseOrder struct {
	ID              int64       `valid:"-"`
	OrderNumber     string      `valid:"required~El número de orden es obligatorio"`
	OrderDate       *utils.Date `valid:"notNil~La fecha de la orden es obligatoria"`
	TrackingCode    string      `valid:"required~El código de seguimiento es obligatorio"`
	BuyerID         int64       `valid:"required~El ID de comprador es obligatorio"`
	ProductRecordID int64       `valid:"required~El ID del registro de producto es obligatorio"`
}

// PurchaseOrderAttributes represents the attributes of a purchase order that can be updated
// Uses pointers to distinguish between null values and default values

type PurchaseOrderAttributes struct {
	OrderNumber     *string `json:"order_number"`
	OrderDate       *string `json:"order_date"`
	TrackingCode    *string `json:"tracking_code"`
	BuyerID         *int64  `json:"buyer_id"`
	ProductRecordID *int64  `json:"product_record_id"`
}

// PurchaseOrderAttributesSwagger is used only for Swagger documentation
// swagger:model
// example: {"order_number":"order#1","order_date":"2021-04-04","tracking_code":"abscf123","buyer_id":1,"product_record_id":1}
type PurchaseOrderAttributesSwagger struct {
	OrderNumber     *string `json:"order_number" example:"order#1"`
	OrderDate       *string `json:"order_date" example:"2021-04-04"`
	TrackingCode    *string `json:"tracking_code" example:"abscf123"`
	BuyerID         *int64  `json:"buyer_id" example:"1"`
	ProductRecordID *int64  `json:"product_record_id" example:"1"`
}

// PurchaseOrderDoc represents the document structure for a purchase order

type PurchaseOrderDoc struct {
	ID int64 `json:"id"`
	PurchaseOrderAttributes
}

// PurchaseOrderResponse represents the response for a single purchase order

type PurchaseOrderResponse struct {
	Data PurchaseOrderDoc `json:"data"`
}

// PurchaseOrdersResponse represents the response for multiple purchase orders

type PurchaseOrdersResponse struct {
	Data []PurchaseOrderDoc `json:"data"`
}

// ParseToPurchaseOrderDoc converts a PurchaseOrder to PurchaseOrderDoc
func (po *PurchaseOrder) ParseToPurchaseOrderDoc() PurchaseOrderDoc {
	var orderDateStr *string
	if po.OrderDate != nil {
		str := po.OrderDate.Time.Format("2006-01-02")
		orderDateStr = &str
	}
	return PurchaseOrderDoc{
		ID: po.ID,
		PurchaseOrderAttributes: PurchaseOrderAttributes{
			OrderNumber:     &po.OrderNumber,
			OrderDate:       orderDateStr,
			TrackingCode:    &po.TrackingCode,
			BuyerID:         &po.BuyerID,
			ProductRecordID: &po.ProductRecordID,
		},
	}
}

// ParseToPurchaseOrder converts PurchaseOrderAttributes to PurchaseOrder
func (poa *PurchaseOrderAttributes) ParseToPurchaseOrder() (po PurchaseOrder, err error) {
	if poa.OrderNumber != nil {
		po.OrderNumber = *poa.OrderNumber
	}
	if poa.OrderDate != nil {
		parsed, parseErr := utils.ParseDate(*poa.OrderDate)
		if parseErr != nil {
			err = parseErr
			return
		}
		po.OrderDate = parsed
	}
	if poa.TrackingCode != nil {
		po.TrackingCode = *poa.TrackingCode
	}
	if poa.BuyerID != nil {
		po.BuyerID = *poa.BuyerID
	}
	if poa.ProductRecordID != nil {
		po.ProductRecordID = *poa.ProductRecordID
	}
	return
}

// ParseToPurchaseOrder converts PurchaseOrderDoc to PurchaseOrder
func (pod *PurchaseOrderDoc) ParseToPurchaseOrder() (PurchaseOrder, error) {
	po, err := pod.PurchaseOrderAttributes.ParseToPurchaseOrder()
	po.ID = pod.ID
	return po, err
}
