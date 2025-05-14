package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config defines the system configuration structure
type Config struct {
	ServerPort        string
	DataPath          string
	APIPrefix         string
	SellerFilePath    string
	ProductFilePath   string
	WarehouseFilePath string
	EmployeeFilePath  string
	BuyerFilePath     string
	SectionFilePath   string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dataPath := getEnvOrDefault("DATA_PATH", "./docs/db")

	// Basic port validation
	port := getEnvOrDefault("SERVER_PORT", "8080")

	config := &Config{
		ServerPort:        port,
		APIPrefix:         getEnvOrDefault("API_PREFIX", "/api/v1"),
		DataPath:          dataPath,
		SellerFilePath:    filepath.Join(dataPath, "seller.json"),
		ProductFilePath:   filepath.Join(dataPath, "product.json"),
		WarehouseFilePath: filepath.Join(dataPath, "warehouse.json"),
		BuyerFilePath:     filepath.Join(dataPath, "buyer.json"),
		EmployeeFilePath:  filepath.Join(dataPath, "employee.json"),
		SectionFilePath:   filepath.Join(dataPath, "section.json"),
	}

	// If no exist, create JSON files
	err := createJSONFiles([]string{
		config.SellerFilePath,
		config.ProductFilePath,
		config.WarehouseFilePath,
		config.BuyerFilePath,
		config.SectionFilePath,
		config.EmployeeFilePath,
	})

	if err != nil {
		return nil, fmt.Errorf("error creating the files: %w", err)
	}

	return config, nil
}

// CreateJSONFiles checks and creates JSON files if they do not exist
func createJSONFiles(paths []string) error {
	for _, path := range paths {
		// Check if the file already exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// Create the file with an empty initial content
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					log.Printf("Error closing file: %v", err)
				}
			}(file)

			// Write an empty array to the file
			encoder := json.NewEncoder(file)
			if err := encoder.Encode([]interface{}{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
