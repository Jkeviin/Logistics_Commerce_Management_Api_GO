package repository

import "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"

type EmployeeRepository interface {
	FindAll() ([]domain.Employee, error)
	Find(id int64) (domain.Employee, error)
	Create(employee domain.Employee) (domain.Employee, error)
	Delete(id int64) error
	Update(id int64, employeePointers domain.Employee) (domain.Employee, error)
	ReportInboundOrders() ([]domain.EmployeeWithInboundOrdersCount, error)
	ReportInboundOrdersById(id int64) (domain.EmployeeWithInboundOrdersCount, error)
}
