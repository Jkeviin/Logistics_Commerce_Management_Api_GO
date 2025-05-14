package application

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func EmployeeRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.EmployeeDoc, int64](cfg.EmployeeFilePath)
	dbDoc, err := ld.LoadToMap(func(v domain.EmployeeDoc) int64 { return v.Id })
	if err != nil {
		panic(err)
	}

	// Convert from WarehouseDoc to Warehouse
	db := make(map[int64]domain.Employee)
	for key, value := range dbDoc {
		db[key] = value.ParseToEmployee()
	}

	// - repository
	rp := repository.NewEmployeeMap(db, ld)
	// - service
	sv := services.NewEmployeeDefault(rp)
	// - handler
	hd := handlers.NewEmployeeDefault(sv)

	return func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.Create())
		r.Delete("/{id}", hd.Delete())
		r.Patch("/{id}", hd.Update())
	}
}
