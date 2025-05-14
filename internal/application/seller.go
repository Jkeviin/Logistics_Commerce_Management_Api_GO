package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SellerRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewSellerSQL(db)
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
