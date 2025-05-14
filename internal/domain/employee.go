package domain

type Employee struct {
	Id           int64  `valid:"-"`
	CardNumberId int64  `valid:"required~El código de tarjeta es obligatorio"`
	FirstName    string `valid:"required~El nombre es obligatorio"`
	LastName     string `valid:"required~El apellido es obligatorio"`
	WarehouseId  int64  `valid:"required~El warehouseId es obligatorio"`
}

type EmployeeAttributes struct {
	CardNumberId int64  `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

type EmployeePatchAttributes struct {
	CardNumberId *int64  `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseId  *int64  `json:"warehouse_id"`
}

type EmployeeDoc struct {
	Id int64 `json:"id"`
	EmployeeAttributes
}

type EmployeeResponse struct {
	Data EmployeeDoc `json:"data"`
}

type EmployeesResponse struct {
	Data []EmployeeDoc `json:"data"`
}

func (w *Employee) ParseToEmployeeDoc() EmployeeDoc {
	return EmployeeDoc{
		Id: w.Id,
		EmployeeAttributes: EmployeeAttributes{
			CardNumberId: w.CardNumberId,
			FirstName:    w.FirstName,
			LastName:     w.LastName,
			WarehouseId:  w.WarehouseId,
		},
	}
}

func (w *EmployeeAttributes) ParseToEmployee() Employee {
	return Employee{
		CardNumberId: w.CardNumberId,
		FirstName:    w.FirstName,
		LastName:     w.LastName,
		WarehouseId:  w.WarehouseId,
	}
}

func (w *EmployeeDoc) ParseToEmployee() Employee {
	employee := w.EmployeeAttributes.ParseToEmployee()
	employee.Id = w.Id
	return employee
}
