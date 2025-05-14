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

func WarehouseRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.WarehouseDoc, int64](cfg.WarehouseFilePath)
	dbDoc, err := ld.LoadToMap(func(v domain.WarehouseDoc) int64 { return v.ID })
	if err != nil {
		panic(err)
	}

	// Convert from WarehouseDoc to Warehouse
	db := make(map[int64]domain.Warehouse)
	for key, value := range dbDoc {
		db[key] = value.ParseToWarehouse()
	}

	// - repository
	rp := repository.NewWarehouseMap(db, ld)
	// - service
	sv := services.NewWarehouseDefault(rp)
	// - handler
	hd := handlers.NewWarehouseDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Delete("/{id}", hd.Delete())
		r.Patch("/{id}", hd.Update())
	}
}
