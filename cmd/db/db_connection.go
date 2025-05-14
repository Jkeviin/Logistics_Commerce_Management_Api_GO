package db

import (
	"database/sql"
	"fmt"

	"github.com/Jkeviin/Logistics_Commerce_Management_Api_GO/cmd/config"
	_ "github.com/go-sql-driver/mysql"
)

// NewSQLConnection establishes and returns a connection to MySQL
func NewSQLConnection(config *config.Config) (*sql.DB, error) {
	// Retrieve environment variables
	user := config.DatabaseUser
	password := config.DatabasePassword
	host := config.DatabaseHost
	port := config.DatabasePort
	database := config.DatabaseName

	// Build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening the connection: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Successfully connected to the database")
	return db, nil
}
