package domain

type Buyer struct {
	Id           int64
	CardNumberId *int    `valid:"required~El número de tarjeta es obligatorio"`
	FirstName    *string `valid:"required~El nombre es obligatorio"`
	LastName     *string `valid:"required~El apellido es obligatorio"`
}

type BuyerAttributes struct {
	CardNumberId *int    `json:"card_number_id"`
	FirstName    *string `json:"first_name" `
	LastName     *string `json:"last_name"`
}

type BuyerDoc struct {
	Id int64 `json:"id"`
	BuyerAttributes
}

type BuyerResponse struct {
	Data BuyerDoc `json:"data"`
}

type BuyerListResponse struct {
	Data []BuyerDoc `json:"data"`
}

// Purchase orders report by buyer
// BuyerPurchaseOrdersReport extends BuyerDoc with purchase orders field

type BuyerPurchaseOrdersReport struct {
	Id              int64 `json:"buyer_id"`
	BuyerAttributes       // Embeds buyer attributes
	PurchaseOrders  int64 `json:"purchase_orders_count"`
}

type BuyerPurchaseOrdersReportResponse struct {
	Data []BuyerPurchaseOrdersReport `json:"data"`
}

func (b *Buyer) ParseToBuyerDoc() BuyerDoc {
	return BuyerDoc{
		Id: b.Id,
		BuyerAttributes: BuyerAttributes{
			CardNumberId: b.CardNumberId,
			FirstName:    b.FirstName,
			LastName:     b.LastName,
		},
	}
}

func (b *BuyerDoc) ParseToBuyer() Buyer {
	buyer := Buyer{}
	buyer.Id = b.Id
	buyer.CardNumberId = b.CardNumberId
	buyer.FirstName = b.FirstName
	buyer.LastName = b.LastName
	return buyer
}

func (b *BuyerAttributes) ParseToBuyer() Buyer {
	buyer := Buyer{}
	buyer.CardNumberId = b.CardNumberId
	buyer.FirstName = b.FirstName
	buyer.LastName = b.LastName
	return buyer
}
