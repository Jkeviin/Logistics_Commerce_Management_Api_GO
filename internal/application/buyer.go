package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func BuyerRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewBuyerSQL(db)
	// - service
	sv := services.NewBuyerDefault(rp)
	// - handler
	hd := handlers.NewBuyerDefault(sv)

	return func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Post("/", hd.Create())
		r.Get("/{id}", hd.GetById())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
		r.Get("/reportPurchaseOrders", hd.ReportPurchaseOrders())
	}
}
