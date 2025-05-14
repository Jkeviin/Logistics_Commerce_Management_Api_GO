package domain

type Section struct {
	ID                 int64    `valid:"-"`
	SectionNumber      *string  `valid:"required~El código del sector es obligatorio"`
	CurrentTemperature *float64 `valid:"notNil~La temperatura actual es obligatoria,float19_2~La temperatura no es compatible con el formato de 19 enteros y 2 decimales"`
	MinimumTemperature *float64 `valid:"notNil~La temperatura minima es obligatoria,float19_2~La temperatura no es compatible con el formato de 19 enteros y 2 decimales"`
	CurrentCapacity    *int     `valid:"notNil~La capacidad actual es obligatoria,range(0|2147483647)~La capacidad actual debe ser mayor o igual a 0"`
	MinimumCapacity    *int     `valid:"notNil~La capacidad mínima es obligatoria,range(1|2147483647)~La capacidad mínima debe ser mayor o igual a 1"`
	MaximumCapacity    *int     `valid:"required~La capacidad máxima es obligatoria,range(1|2147483647)~La capacidad máxima debe ser mayor o igual a 1"`
	WarehouseID        *int64   `valid:"required~El id del warehouse asociado es obligatorio"`
	ProductTypeID      *int     `valid:"required~El id del producto es obligatorio"`
}

type SectionAttributes struct {
	SectionNumber      *string  `json:"section_number"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	CurrentCapacity    *int     `json:"current_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	WarehouseID        *int64   `json:"warehouse_id"`
	ProductTypeID      *int     `json:"product_type_id"`
}

type SectionDoc struct {
	ID int64 `json:"id"`
	SectionAttributes
}

type SectionResponse struct {
	Data SectionDoc `json:"data"`
}

type SectionsResponse struct {
	Data []SectionDoc `json:"data"`
}

type SectionWithProductCount struct {
	SectionID     int64  `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductCount  int    `json:"product_count"`
}

type SectionWithProductsCountResponse struct {
	Data []SectionWithProductCount `json:"data"`
}

func (s *SectionAttributes) ParseToSection() Section {
	return Section{
		SectionNumber:      s.SectionNumber,
		CurrentTemperature: s.CurrentTemperature,
		MinimumTemperature: s.MinimumTemperature,
		CurrentCapacity:    s.CurrentCapacity,
		MinimumCapacity:    s.MinimumCapacity,
		MaximumCapacity:    s.MaximumCapacity,
		WarehouseID:        s.WarehouseID,
		ProductTypeID:      s.ProductTypeID,
		//ProductBatches:     s.ProductBatches,
	}
}

func (s *SectionDoc) ParseToSection() Section {
	return Section{
		ID:                 s.ID,
		SectionNumber:      s.SectionNumber,
		CurrentTemperature: s.CurrentTemperature,
		MinimumTemperature: s.MinimumTemperature,
		CurrentCapacity:    s.CurrentCapacity,
		MinimumCapacity:    s.MinimumCapacity,
		MaximumCapacity:    s.MaximumCapacity,
		WarehouseID:        s.WarehouseID,
		ProductTypeID:      s.ProductTypeID,
		//ProductBatches:     *s.ProductBatches,
	}
}

func (s *Section) ParseToSectionDoc() SectionDoc {
	return SectionDoc{
		ID: s.ID,
		SectionAttributes: SectionAttributes{
			SectionNumber:      s.SectionNumber,
			CurrentTemperature: s.CurrentTemperature,
			MinimumTemperature: s.MinimumTemperature,
			CurrentCapacity:    s.CurrentCapacity,
			MinimumCapacity:    s.MinimumCapacity,
			MaximumCapacity:    s.MaximumCapacity,
			WarehouseID:        s.WarehouseID,
			ProductTypeID:      s.ProductTypeID,
			//ProductBatches:     s.ProductBatches,
		},
	}
}
