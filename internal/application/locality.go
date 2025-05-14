package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func LocalityRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewLocalitySQL(db)
	// - service
	sv := services.NewLocalityDefault(rp)
	// - handler
	hd := handlers.NewLocalityDefault(sv)

	return func(r chi.Router) {
		r.Get("/reportSellers", hd.FindSellers())
		r.Post("/", hd.Create())
		r.Get("/reportCarries", hd.GetCantCarriesPerLocality())
	}
}
