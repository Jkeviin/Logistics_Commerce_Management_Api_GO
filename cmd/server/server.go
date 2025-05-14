package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/db"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	_ "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/docs" // Import generate documents for swaggo
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/application"
	customMiddleware "github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// ServerChi is the HTTP server implementation using Chi
type ServerChi struct {
	config *config.Config
	router *chi.Mux
}

// NewServerChi creates a new server instance
func NewServerChi(cfg *config.Config) *ServerChi {
	s := &ServerChi{
		config: cfg,
		router: chi.NewRouter(),
	}
	// Initialize the database connection
	dbConnection, err := db.NewSQLConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}
	s.setupRouter(dbConnection)
	return s
}

// setupRouter configures routes and middlewares
func (s *ServerChi) setupRouter(db *sql.DB) {
	// Swagger route without middlewares
	s.router.Get("/swagger/*", httpSwagger.WrapHandler)

	// Route group with common middlewares
	s.router.Group(func(r chi.Router) {
		// Middlewares
		r.Use(customMiddleware.Logger)
		r.Use(chiMiddleware.Recoverer)

		// Configure base prefix for all API routes
		r.Route(s.config.APIPrefix, func(r chi.Router) {
			// API routes
			r.Route("/sellers", application.SellerRoutes(db))
			r.Route("/warehouses", application.WarehouseRoutes(db))
			r.Route("/employees", application.EmployeeRoutes(db))
			r.Route("/buyers", application.BuyerRoutes(db))
			r.Route("/products", application.ProductRoutes(db))
			r.Route("/sections", application.SectionRoutes(db))
			r.Route("/purchaseOrders", application.PurchaseOrderRoutes(db))
			r.Route("/productBatches", application.ProductBatchRoutes(db))
			r.Route("/localities", application.LocalityRoutes(db))
			r.Route("/carries", application.CarriesRoutes(db))
			r.Route("/inboundOrders", application.InboundOrdersRoutes(db))
			r.Route("/productRecords", application.ProductRecordRoutes(db))
		})
	})
}

// Run starts the HTTP server
func (s *ServerChi) Run() error {
	serverPort := s.config.ServerPort
	if serverPort[0] != ':' {
		serverPort = ":" + serverPort
	}

	fmt.Printf("Starting server on port %s\n", serverPort)
	fmt.Printf("Swagger documentation available at: http://localhost%s/swagger/index.html\n", serverPort)

	return http.ListenAndServe(serverPort, s.router)
}
