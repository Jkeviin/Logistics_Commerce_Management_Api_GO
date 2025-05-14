package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/go-chi/chi/v5"
)

func ProductBatchRoutes(db *sql.DB) func(route chi.Router) {
	// - repository
	rp := repository.NewProductBatchSQL(db)
	// - service
	sv := services.NewProductBatchDefault(rp)
	// - handler
	hd := handlers.NewProductBatchDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
	}
}
