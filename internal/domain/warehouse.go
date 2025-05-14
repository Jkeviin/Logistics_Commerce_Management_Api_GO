package domain

type Warehouse struct {
	ID                 int64    `valid:"-"`
	WarehouseCode      *string  `valid:"required~El código del warehouse es obligatorio"`
	Address            *string  `valid:"required~La dirección es obligatoria"`
	Telephone          *string  `valid:"required~El teléfono es obligatorio,length(7|15)~El teléfono debe contener entre 7 y 15 dígitos"`
	MinimumCapacity    *int     `valid:"required~La capacidad mínima es obligatoria,range(1|2147483647)~La capacidad mínima debe ser mayor o igual a 1"`
	MinimumTemperature *float64 `valid:"notNil~La temperatura mínima es obligatoria,float19_2~La temperatura no es compatible con el formato de 19 enteros y 2 decimales"`
	LocalityID         *string  `valid:"required~El ID de la localidad es obligatorio"`
}

type WarehouseAttributes struct {
	WarehouseCode      *string  `json:"warehouse_code"`
	Address            *string  `json:"address"`
	Telephone          *string  `json:"telephone"`
	MinimumCapacity    *int     `json:"minimun_capacity"`
	MinimumTemperature *float64 `json:"minimun_temperature"`
	LocalityID         *string  `json:"locality_id"`
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
			WarehouseCode:      w.WarehouseCode,
			Address:            w.Address,
			Telephone:          w.Telephone,
			MinimumCapacity:    w.MinimumCapacity,
			MinimumTemperature: w.MinimumTemperature,
			LocalityID:         w.LocalityID,
		},
	}
}

func (w *WarehouseDoc) ParseToWarehouse() Warehouse {
	warehouse := w.WarehouseAttributes.ParseToWarehouse()
	warehouse.ID = w.ID
	return warehouse
}

func (w *WarehouseAttributes) ParseToWarehouse() Warehouse {
	return Warehouse{
		WarehouseCode:      w.WarehouseCode,
		Address:            w.Address,
		Telephone:          w.Telephone,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
		LocalityID:         w.LocalityID,
	}
}
