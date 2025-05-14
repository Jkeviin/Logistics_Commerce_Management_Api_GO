package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func EmployeeRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewEmployeeSQL(db)
	// - service
	sv := services.NewEmployeeDefault(rp)
	// - handler
	hd := handlers.NewEmployeeDefault(sv)

	return func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.Create())
		r.Delete("/{id}", hd.Delete())
		r.Patch("/{id}", hd.Update())

		// - GET /employees/reportInboundOrders?id=
		r.Get("/reportInboundOrders", hd.ReportInboundOrders())
	}
}
