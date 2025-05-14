package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/go-chi/chi/v5"
)

func InboundOrdersRoutes(db *sql.DB) func(route chi.Router) {
	// - repository
	rp := repository.NewInboundOrdersSQL(db)
	// - service
	sv := services.NewInboundOrdersDefault(rp)
	// - handler
	hd := handlers.NewInboundOrdersDefault(sv)

	return func(r chi.Router) {
		r.Post("/", hd.Create())
	}
}
