package domain

type Buyer struct {
	Id           int64  `valid:"-"`
	CardNumberId int    `valid:"required~El número de tarjeta es obligatorio"`
	FirstName    string `valid:"required~El nombre es obligatorio"`
	LastName     string `valid:"required~El apellido es obligatorio"`
}

type BuyerAttributes struct {
	CardNumberId *int    `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
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

func (b *Buyer) ParseToBuyerDoc() BuyerDoc {
	return BuyerDoc{
		Id: b.Id,
		BuyerAttributes: BuyerAttributes{
			CardNumberId: &b.CardNumberId,
			FirstName:    &b.FirstName,
			LastName:     &b.LastName,
		},
	}
}

func (b *BuyerDoc) ParseToBuyer() Buyer {
	buyer := Buyer{}
	buyer.Id = b.Id

	if b.CardNumberId != nil {
		buyer.CardNumberId = *b.CardNumberId
	}
	if b.FirstName != nil {
		buyer.FirstName = *b.FirstName
	}
	if b.LastName != nil {
		buyer.LastName = *b.LastName
	}
	return buyer
}

func (b *BuyerAttributes) ParseToBuyer() Buyer {
	buyer := Buyer{}
	if b.CardNumberId != nil {
		buyer.CardNumberId = *b.CardNumberId
	}
	if b.FirstName != nil {
		buyer.FirstName = *b.FirstName
	}
	if b.LastName != nil {
		buyer.LastName = *b.LastName
	}
	return buyer
}
