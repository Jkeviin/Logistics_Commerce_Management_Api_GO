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

func BuyerRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.BuyerDoc, int64](cfg.BuyerFilePath)
	db, err := ld.LoadToMap(func(v domain.BuyerDoc) int64 { return v.Id })
	if err != nil {
		panic(err)
	}

	dbBuyer := make(map[int64]domain.Buyer)
	for _, v := range db {
		dbBuyer[v.Id] = v.ParseToBuyer()
	}

	// - repository
	rp := repository.NewBuyerMap(dbBuyer, ld)
	// - service
	sv := services.NewBuyerDefault(rp)
	// - handler
	hd := handlers.NewBuyerDefault(sv)

	return func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Post("/", hd.Create())
		r.Get("/{id}", hd.GetById())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	}
}
