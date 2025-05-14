package main

import (
	"log"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/server"
)

//	@title			API de Sellers y Warehouses
//	@version		1.0
//	@description	API para gestionar sellers y warehouses
//	@BasePath		/api/v1
//	@host			localhost:8080

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize and run server
	app := server.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
