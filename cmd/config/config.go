package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

// Config defines the system configuration structure
type Config struct {
	ServerPort       string
	APIPrefix        string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Basic port validation
	port := getEnvOrDefault("SERVER_PORT", "8080")

	config := &Config{
		ServerPort:       port,
		APIPrefix:        getEnvOrDefault("API_PREFIX", "/api/v1"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
	}

	values := reflect.ValueOf(*config)
	keys := reflect.TypeOf(*config)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).IsZero() {
			return nil, fmt.Errorf("missing required configuration: %s", keys.Field(i).Name)
		}
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
