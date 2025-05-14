package filters

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

// SellerFilter represents the possible filters for searching sellers
type SellerFilter struct {
	ID          *int64  `json:"id,omitempty"`
	CID         *int    `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
}

// HasFilters checks if the filter has any defined fields
func (f *SellerFilter) HasFilters() bool {
	return utils.HasFilters(f)
}

// Matches checks if a seller matches the specified filters
func (f *SellerFilter) Matches(seller domain.Seller) bool {
	return utils.Matches(f, seller)
}
