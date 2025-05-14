package repository

import (
	"time"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/utils"
)

func NewEmployeeMap(db map[int64]domain.Employee, loader *loader.JSONLoader[domain.EmployeeDoc, int64]) *EmployeeMap {
	// default db
	defaultDb := make(map[int64]domain.Employee)
	if db != nil {
		defaultDb = db
	}
	return &EmployeeMap{
		db:     defaultDb,
		loader: loader,
	}
}

type EmployeeMap struct {
	db     map[int64]domain.Employee
	loader *loader.JSONLoader[domain.EmployeeDoc, int64]
}

func (r *EmployeeMap) FindAll() ([]domain.Employee, error) {

	employees := utils.MapToSlice(r.db)

	return employees, nil
}

func (r *EmployeeMap) Find(id int64) (domain.Employee, error) {
	employee, ok := r.db[id]
	if !ok {
		return domain.Employee{}, utils.ErrEmployeeNotFound
	}
	return employee, nil
}

func (r *EmployeeMap) Create(employee domain.Employee) (domain.Employee, error) {
	employee.Id = time.Now().UnixNano()
	_, ok := r.db[employee.Id]
	if ok {
		return domain.Employee{}, utils.ErrEmployeeAlreadyExists
	}
	r.db[employee.Id] = employee

	// Save to JSON
	if err := r.loader.SaveToJSON(r.getDBInEmployeeDoc()); err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeMap) FindByCardNum(number int64) (domain.Employee, error) {
	for _, value := range r.db {
		if value.CardNumberId == number {
			return value, nil
		}
	}
	return domain.Employee{}, utils.ErrEmployeeNotFound
}

func (r *EmployeeMap) getDBInEmployeeDoc() map[int64]domain.EmployeeDoc {
	dbDoc := make(map[int64]domain.EmployeeDoc)
	for key, value := range r.db {
		dbDoc[key] = value.ParseToEmployeeDoc()
	}
	return dbDoc
}

func (r *EmployeeMap) Delete(id int64) error {
	_, ok := r.db[id]
	if !ok {
		return utils.ErrEmployeeNotFound
	}
	delete(r.db, id)

	// Persist changes to JSON after deletion
	if err := r.loader.SaveToJSON(r.getDBInEmployeeDoc()); err != nil {
		return err
	}
	return nil
}

func (r *EmployeeMap) Update(id int64, employee domain.Employee) (domain.Employee, error) {
	_, ok := r.db[id]
	if !ok {
		return domain.Employee{}, utils.ErrEmployeeNotFound
	}
	employee.Id = id
	r.db[id] = employee

	// Persist changes to JSON after update
	if err := r.loader.SaveToJSON(r.getDBInEmployeeDoc()); err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}
