package domain

type Warehouse struct {
	ID                 int64  `valid:"-"`
	WarehouseCode      string `valid:"required~El código del warehouse es obligatorio"`
	Address            string `valid:"required~La dirección es obligatoria"`
	Telephone          string `valid:"required~El teléfono es obligatorio,length(7|15)~El teléfono debe contener entre 7 y 15 dígitos"`
	MinimumCapacity    int    `valid:"required~La capacidad mínima es obligatoria,range(1|2147483647)~La capacidad mínima debe ser mayor o igual a 1"`
	MinimumTemperature int    `valid:"-"`
}

type WarehouseAttributes struct {
	WarehouseCode      *string `json:"warehouse_code"`
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	MinimumCapacity    *int    `json:"minimun_capacity"`
	MinimumTemperature *int    `json:"minimun_temperature"`
}

type WarehouseDoc struct {
	ID int64 `json:"id"`
	WarehouseAttributes
}

type WarehouseResponse struct {
	Data WarehouseDoc `json:"data"`
}

type WarehousesResponse struct {
	Data []WarehouseDoc `json:"data"`
}

func (w *Warehouse) ParseToWarehouseDoc() WarehouseDoc {
	return WarehouseDoc{
		ID: w.ID,
		WarehouseAttributes: WarehouseAttributes{
			WarehouseCode:      &w.WarehouseCode,
			Address:            &w.Address,
			Telephone:          &w.Telephone,
			MinimumCapacity:    &w.MinimumCapacity,
			MinimumTemperature: &w.MinimumTemperature,
		},
	}
}

func (w *WarehouseDoc) ParseToWarehouse() Warehouse {
	warehouse := w.WarehouseAttributes.ParseToWarehouse()
	warehouse.ID = w.ID
	return warehouse
}

func (w *WarehouseAttributes) ParseToWarehouse() (warehouse Warehouse) {
	if w.WarehouseCode != nil {
		warehouse.WarehouseCode = *w.WarehouseCode
	}
	if w.Address != nil {
		warehouse.Address = *w.Address
	}
	if w.Telephone != nil {
		warehouse.Telephone = *w.Telephone
	}
	if w.MinimumCapacity != nil {
		warehouse.MinimumCapacity = *w.MinimumCapacity
	}
	if w.MinimumTemperature != nil {
		warehouse.MinimumTemperature = *w.MinimumTemperature
	}

	return
}
