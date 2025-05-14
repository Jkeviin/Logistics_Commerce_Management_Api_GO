package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/go-chi/chi/v5"
)

func PurchaseOrderRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		poRepo := repository.NewPurchaseOrderSQL(db)
		service := services.NewPurchaseOrderDefault(poRepo)
		handler := handlers.NewPurchaseOrderDefault(service)

		r.Post("/", handler.Create())
	}
}
