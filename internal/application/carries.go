package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func CarriesRoutes(db *sql.DB) func(chi.Router) {
	rp := repository.NewCarriesSQL(db)
	// - service
	sv := services.NewCarriesDefault(rp)
	// - handler
	hd := handlers.NewCarriesDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
	}
}
