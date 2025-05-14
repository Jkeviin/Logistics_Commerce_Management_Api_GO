package services

import (
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/interfaces/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewEmployeeDefault(rp repository.EmployeeRepository) *EmployeeDefault {
	return &EmployeeDefault{rp: rp}
}

type EmployeeDefault struct {
	rp repository.EmployeeRepository
}

func (s *EmployeeDefault) FindAll() ([]domain.Employee, error) {
	return s.rp.FindAll()
}

func (s *EmployeeDefault) Find(id int64) (domain.Employee, error) {
	return s.rp.Find(id)
}

func (s *EmployeeDefault) Create(employee domain.Employee) (domain.Employee, error) {
	err := utils.ValidateStructGoValidator(employee)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("%w: %w", utils.ErrValidation, err)
	}

	return s.rp.Create(employee)
}

func (s *EmployeeDefault) Delete(id int64) error {
	return s.rp.Delete(id)
}

func (s *EmployeeDefault) Update(id int64, employeePointer domain.EmployeeAttributes) (domain.Employee, error) {
	employeeToUpdate, err := s.rp.Find(id)
	if err != nil {
		return domain.Employee{}, utils.ErrEmployeeNotFound
	}
	utils.UpdateStruct(&employeeToUpdate, &employeePointer)

	return s.rp.Update(id, employeeToUpdate)
}

func (sv *EmployeeDefault) ReportInboundOrdersByEmployee(idEmployee int64) ([]domain.EmployeeWithInboundOrdersCount, error) {
	if idEmployee == 0 {
		return sv.rp.ReportInboundOrders()
	}
	empl, err := sv.rp.ReportInboundOrdersById(idEmployee)
	if err != nil {
		return nil, err
	}
	return []domain.EmployeeWithInboundOrdersCount{empl}, nil
}
