package application

import (
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/domain"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/handlers"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/loader"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/repository"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/services"
	"github.com/go-chi/chi/v5"
)

func SectionRoutes(cfg *config.Config) func(chi.Router) {
	// - loader
	ld := loader.NewJSONLoader[domain.SectionDoc, int64](cfg.SectionFilePath)
	dbDoc, err := ld.LoadToMap(func(s domain.SectionDoc) int64 { return s.ID })
	if err != nil {
		panic(err)
	}

	// Convert from SectionDoc to Section
	db := make(map[int64]domain.Section)
	for key, value := range dbDoc {
		db[key] = value.ParseToSection()
	}

	// - repository
	rp := repository.NewSectionMap(db, ld)
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
	}

}
