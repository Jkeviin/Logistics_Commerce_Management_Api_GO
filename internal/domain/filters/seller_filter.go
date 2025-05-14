package filters

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// SellerFilter represents the possible filters for searching sellers
type SellerFilter struct {
	ID          *int64  `json:"id,omitempty"`
	CID         *string `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
	LocalityID  *int    `json:"locality_id,omitempty"`
}

// HasFilters checks if the filter has any defined fields
func (f *SellerFilter) HasFilters() bool {
	return utils.HasFilters(f)
}

// FieldColumnMap returns the mapping between the field names and the SQL columns
func (f *SellerFilter) FieldColumnMap() map[string]string {
	return map[string]string{
		"ID":          "id",
		"CID":         "cid",
		"CompanyName": "company_name",
		"Address":     "address",
		"Telephone":   "telephone",
		"LocalityID":  "locality_id",
	}
}
