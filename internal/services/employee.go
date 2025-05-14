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

	employee, err := s.rp.FindAll()
	if err != nil {
		return nil, utils.ErrInternalServer
	}

	return employee, nil
}

func (s *EmployeeDefault) Find(id int64) (domain.Employee, error) {
	result, err := s.rp.Find(id)
	if err != nil {
		return domain.Employee{}, utils.ErrEmployeeNotFound
	}

	return result, nil
}

func (s *EmployeeDefault) Create(employee domain.Employee) (domain.Employee, error) {
	err := utils.ValidateStructGoValidator(employee)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("Error al validar el employee: %w", err)
	}

	_, err = s.rp.FindByCardNum(employee.CardNumberId)
	if err == nil {
		return domain.Employee{}, utils.ErrCardNumberAlreadyExists
	}

	result, err := s.rp.Create(employee)
	if err != nil {
		return domain.Employee{}, utils.ErrInternalServer
	}

	return result, nil
}

func (s *EmployeeDefault) Delete(id int64) error {
	_, err := s.rp.Find(id)
	if err != nil {
		return utils.ErrEmployeeNotFound
	}

	err = s.rp.Delete(id)
	if err != nil {
		return utils.ErrInternalServer
	}

	return nil
}

func (s *EmployeeDefault) Update(id int64, employeePointer domain.EmployeePatchAttributes) (domain.Employee, error) {
	employeeToUpdate, err := s.rp.Find(id)
	if err != nil {
		return domain.Employee{}, utils.ErrEmployeeNotFound
	}

	utils.UpdateStruct(&employeeToUpdate, &employeePointer)

	if err := utils.ValidateStructGoValidator(employeeToUpdate); err != nil {
		return domain.Employee{}, fmt.Errorf("Error al validar el empleado: %w", err)
	}

	employedSearched, err := s.rp.FindByCardNum(employeeToUpdate.CardNumberId)
	if err == nil && employedSearched.Id != id {
		return domain.Employee{}, utils.ErrCardNumberAlreadyExists
	}

	result, err := s.rp.Update(id, employeeToUpdate)
	if err != nil {
		return domain.Employee{}, utils.ErrInternalServer
	}

	return result, nil
}
