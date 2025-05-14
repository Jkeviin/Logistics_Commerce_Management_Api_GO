package application

import (
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.Product, int64](cfg.ProductFilePath)
	db, err := ld.LoadToMap(func(v domain.Product) int64 { return int64(v.Id) }) // Assuming ID is the correct field
	if err != nil {
		panic(err)
	}
	rp := repository.NewProductMap(db, ld)
	// - service
	sv := services.NewProductDefault(rp)
	// - handler
	hd := handlers.NewProductDefault(sv)

	return func(r chi.Router) {
		r.Get("/", http.HandlerFunc(hd.GetAll()))
		r.Post("/", http.HandlerFunc(hd.Create()))
		r.Get("/{id}", http.HandlerFunc(hd.GetByID()))
		r.Patch("/{id}", http.HandlerFunc(hd.Update()))
		r.Delete("/{id}", http.HandlerFunc(hd.Delete()))
	}
}
