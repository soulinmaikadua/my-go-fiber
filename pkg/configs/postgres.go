package configs

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
)

func PostgresConnection() (*sqlx.DB, error) {
	// Define database connection for PostgreSQL
	fmt.Println("DB_SERVER_URL:", os.Getenv("DB_SERVER_URL"))
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connect to database: %v", err)
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // Close connection
		return nil, fmt.Errorf("error, not sent ping to database: %v", err)
	}
	fmt.Println("Connection to database")
	return db, nil
}
