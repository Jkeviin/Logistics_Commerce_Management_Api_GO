package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(db *sql.DB) func(chi.Router) {
	rp := repository.NewProductSQL(db)
	// - service
	sv := services.NewProductDefault(rp)
	// - handler
	hd := handlers.NewProductDefault(sv)

	return func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Post("/", hd.Create())
		r.Get("/{id}", hd.GetByID())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
		r.Get("/reportRecords", hd.GetReportRecords())
	}
}
