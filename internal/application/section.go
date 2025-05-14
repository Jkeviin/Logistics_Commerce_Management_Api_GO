package application

import (
	"database/sql"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/go-chi/chi/v5"
)

func SectionRoutes(db *sql.DB) func(chi.Router) {
	// - repository
	rp := repository.NewSectionSQL(db)
	// - service
	sv := services.NewSectionDefault(rp)
	// - handler
	hd := handlers.NewSectionDefault(sv)

	return func(r chi.Router) {
		// - GET /sections
		r.Get("/", hd.FindAll())
		r.Get("/{id}", hd.FindById())

		// - POST /sections
		r.Post("/", hd.Create())

		// - PATCH /sections/{id}
		r.Patch("/{id}", hd.Update())

		// - DELETE /sections/{id}
		r.Delete("/{id}", hd.Delete())

		// - GET /sections/reportProducts?id=
		r.Get("/reportProducts", hd.ReportProducts())
	}

}
