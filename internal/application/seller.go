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

func SellerRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.SellerDoc, int64](cfg.SellerFilePath)
	dbDoc, err := ld.LoadToMap(func(v domain.SellerDoc) int64 { return v.ID })
	if err != nil {
		panic(err)
	}

	// Convert from SellerDoc to Seller
	db := make(map[int64]domain.Seller)
	for key, value := range dbDoc {
		seller := value.ParseToSeller()
		seller.ID = key // Asegurar que el ID se mantiene
		db[key] = seller
	}

	// - repository
	rp := repository.NewSellerMap(db, ld)
	// - service
	sv := services.NewSellerDefault(rp)
	// - handler
	hd := handlers.NewSellerDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Delete("/{id}", hd.Delete())
		r.Patch("/{id}", hd.Update())
	}
}
