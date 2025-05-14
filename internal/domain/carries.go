package domain

// Carry represents a carry in the system.
type CarriesAttributes struct {
	CID         *string `json:"cid" valid:"required~El número de Carry es obligatorio"`
	CompanyName *string `json:"company_name" valid:"required~El nombre de la compañía es obligatorio"`
	Address     *string `json:"address" valid:"required~La dirección es obligatoria"`
	Telephone   *string `json:"telephone" valid:"required~El teléfono es obligatorio"`
	LocalityID  *string `json:"locality_id" valid:"required~El ID de la localidad es obligatorio"`
}

type Carries struct {
	ID int64 `json:"id" valid:"-"` // ID is not required for creation
	CarriesAttributes
}

// CarryDoc represents a carry document in the system.
type CarriesDoc struct {
	ID int64 `json:"id"`
	CarriesAttributes
}

// CarriesResponse represents the response structure for multiple carries.
type CarriesResponse struct {
	Data []CarriesDoc `json:"data"`
}

// CarryResponse represents the response structure for a single carry.
type CarryResponse struct {
	Data CarriesDoc `json:"data"`
}

// ParseToCarryDoc converts a Carry to a CarryDoc.
func (c *Carries) ParseToCarrieDoc() CarriesDoc {
	return CarriesDoc{
		ID:                c.ID,
		CarriesAttributes: c.CarriesAttributes,
	}
}

// ParseToCarry converts a CarryDoc to a Carry.
func (c *CarriesDoc) ParseToCarrie() Carries {
	return Carries{
		ID:                c.ID,
		CarriesAttributes: c.CarriesAttributes,
	}
}
