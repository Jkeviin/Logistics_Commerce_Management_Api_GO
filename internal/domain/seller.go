package domain

// Seller represents a seller in the domain
type Seller struct {
	ID          int64  `valid:"-"`
	CID         string `valid:"required~El CID es obligatorio"`
	CompanyName string `valid:"required~La razón social es obligatoria"`
	Address     string `valid:"required~La dirección es obligatoria"`
	Telephone   string `valid:"required~El teléfono es obligatorio,length(7|15)~El teléfono debe contener entre 7 y 15 dígitos"`
	LocalityID  string `valid:"required~El ID de localidad es obligatorio"`
}

// SellerAttributes represents the attributes of a seller that can be updated
type SellerAttributes struct {
	CID         *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
	LocalityID  *string `json:"locality_id"`
}

// SellerDoc represents the document structure for a seller
type SellerDoc struct {
	ID int64 `json:"id"`
	SellerAttributes
}

// SellerResponse represents the response structure for a single seller
type SellerResponse struct {
	Data SellerDoc `json:"data"`
}

// SellersResponse represents the response structure for multiple sellers
type SellersResponse struct {
	Data []SellerDoc `json:"data"`
}

// ParseToSellerDoc converts a Seller to SellerDoc
func (s *Seller) ParseToSellerDoc() SellerDoc {
	return SellerDoc{
		ID: s.ID,
		SellerAttributes: SellerAttributes{
			CID:         &s.CID,
			CompanyName: &s.CompanyName,
			Address:     &s.Address,
			Telephone:   &s.Telephone,
			LocalityID:  &s.LocalityID,
		},
	}
}

// ParseToSeller converts SellerAttributes to Seller
func (s *SellerAttributes) ParseToSeller() (seller Seller) {
	if s.CID != nil {
		seller.CID = *s.CID
	}
	if s.CompanyName != nil {
		seller.CompanyName = *s.CompanyName
	}
	if s.Address != nil {
		seller.Address = *s.Address
	}
	if s.Telephone != nil {
		seller.Telephone = *s.Telephone
	}
	if s.LocalityID != nil {
		seller.LocalityID = *s.LocalityID
	}
	return
}

// ParseToSeller converts SellerDoc to Seller
func (s *SellerDoc) ParseToSeller() Seller {
	seller := s.SellerAttributes.ParseToSeller()
	seller.ID = s.ID
	return seller
}
