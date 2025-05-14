package domain

type Section struct {
	ID                 int64  `valid:"-"`
	SectionNumber      string `valid:"required~El código del sector es obligatorio"`
	CurrentTemperature int    `valid:"int~La temperatura actual es obligatoria"`
	MinimumTemperature int    `valid:"int~La temperatura minima es obligatoria"`
	CurrentCapacity    int    `valid:"int~La capacidad actual es obligatoria,range(0|2147483647)~La capacidad actual debe ser mayor o igual a 0"`
	MinimumCapacity    int    `valid:"int~La capacidad mínima es obligatoria,range(1|2147483647)~La capacidad mínima debe ser mayor o igual a 1"`
	MaximumCapacity    int    `valid:"required~La capacidad máxima es obligatoria,range(1|2147483647)~La capacidad máxima debe ser mayor o igual a 1"`
	WarehouseID        int64  `valid:"required~El id del warehouse asociado es obligatorio"`
	ProductTypeID      int    `valid:"required~El id del producto es obligatorio"`
	//ProductBatches     []ProductBatch `json:"product_batches"`
}

type SectionAttributes struct {
	SectionNumber      string `json:"section_number"`
	CurrentTemperature *int   `json:"current_temperature"`
	MinimumTemperature *int   `json:"minimum_temperature"`
	CurrentCapacity    *int   `json:"current_capacity"`
	MinimumCapacity    *int   `json:"minimum_capacity"`
	MaximumCapacity    int    `json:"maximum_capacity"`
	WarehouseID        int64  `json:"warehouse_id"`
	ProductTypeID      int    `json:"product_type_id"`
	//ProductBatches     []ProductBatch `json:"product_batches"`
}

type SectionPatchAttributes struct {
	SectionNumber      *string `json:"section_number"`
	CurrentTemperature *int    `json:"current_temperature"`
	MinimumTemperature *int    `json:"minimum_temperature"`
	CurrentCapacity    *int    `json:"current_capacity"`
	MinimumCapacity    *int    `json:"minimum_capacity"`
	MaximumCapacity    *int    `json:"maximum_capacity"`
	WarehouseID        *int64  `json:"warehouse_id"`
	ProductTypeID      *int    `json:"product_type_id"`
	//ProductBatches     *[]ProductBatch `json:"product_batches"`
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

func (s *SectionAttributes) ParseToSection() Section {
	return Section{
		SectionNumber:      s.SectionNumber,
		CurrentTemperature: *s.CurrentTemperature,
		MinimumTemperature: *s.MinimumTemperature,
		CurrentCapacity:    *s.CurrentCapacity,
		MinimumCapacity:    *s.MinimumCapacity,
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
		CurrentTemperature: *s.CurrentTemperature,
		MinimumTemperature: *s.MinimumTemperature,
		CurrentCapacity:    *s.CurrentCapacity,
		MinimumCapacity:    *s.MinimumCapacity,
		MaximumCapacity:    s.MaximumCapacity,
		WarehouseID:        s.WarehouseID,
		ProductTypeID:      s.ProductTypeID,
		//ProductBatches:     s.ProductBatches,
	}
}

func (s *Section) ParseToSectionDoc() SectionDoc {
	return SectionDoc{
		ID: s.ID,
		SectionAttributes: SectionAttributes{
			SectionNumber:      s.SectionNumber,
			CurrentTemperature: &s.CurrentTemperature,
			MinimumTemperature: &s.MinimumTemperature,
			CurrentCapacity:    &s.CurrentCapacity,
			MinimumCapacity:    &s.MinimumCapacity,
			MaximumCapacity:    s.MaximumCapacity,
			WarehouseID:        s.WarehouseID,
			ProductTypeID:      s.ProductTypeID,
			//ProductBatches:     s.ProductBatches,
		},
	}
}
