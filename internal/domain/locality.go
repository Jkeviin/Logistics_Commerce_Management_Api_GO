package domain

type Locality struct {
	Id           string  `json:"id" valid:"required~El id es obligatorio"`
	LocalityName *string `json:"locality_name" valid:"required~El nombre de la localidad es obligatorio"`
	ProvinceName *string `json:"province_name" valid:"required~El nombre de la provincia es obligatorio"`
	CountryName  *string `json:"country_name" valid:"required~El nombre del país es obligatorio"`
}

type LocalityAttributes struct {
	Id           string  `json:"id"`
	LocalityName *string `json:"locality_name"`
	ProvinceName *string `json:"province_name"`
	CountryName  *string `json:"country_name"`
}

type LocalityWithCount struct {
	Id           string  `json:"locality_id"`
	LocalityName *string `json:"locality_name"`
	CantCarries  *int    `json:"carries_count"`
}

type LocalityWithCountResponse struct {
	Data []LocalityWithCount `json:"data"`
}

type LocalityDoc struct {
	Id string `json:"id"`
	LocalityAttributes
}

type LocalityResponse struct {
	Data LocalityDoc `json:"data"`
}

type LocalitiesResponse struct {
	Data []LocalityDoc `json:"data"`
}

type LocalityReportResponse struct {
	Data []LocalityReportAttributes `json:"data"`
}
type LocalityReportAttributes struct {
	LocalityID   string `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}

// ParseToLocalityDoc convierte una entidad Locality a LocalityDoc
func (w *Locality) ParseToLocalityDoc() LocalityDoc {
	return LocalityDoc{
		Id: w.Id,
		LocalityAttributes: LocalityAttributes{
			Id:           w.Id,
			LocalityName: w.LocalityName,
			ProvinceName: w.ProvinceName,
			CountryName:  w.CountryName,
		},
	}
}

// ParseToLocality convierte LocalityAttributes a una entidad Locality
func (w *LocalityAttributes) ParseToLocality() Locality {
	return Locality{
		Id:           w.Id,
		LocalityName: w.LocalityName,
		ProvinceName: w.ProvinceName,
		CountryName:  w.CountryName,
	}
}

// ParseToLocality convierte LocalityDoc a una entidad Locality
func (w *LocalityDoc) ParseToLocality() Locality {
	return Locality{
		Id:           w.Id,
		LocalityName: w.LocalityName,
		ProvinceName: w.ProvinceName,
		CountryName:  w.CountryName,
	}
}
