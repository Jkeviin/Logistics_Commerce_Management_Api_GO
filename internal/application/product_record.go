package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func ProductRecordRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewProductRecordSQL(db)
	// - service
	sv := services.NewProductRecordDefault(rp)
	// - handler
	hd := handlers.NewProductRecordDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
	}
}
