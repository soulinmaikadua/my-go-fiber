package configs

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
)

func PostgresConnection() (*sqlx.DB, error) {
	fmt.Println("PostgresConnection")
	// Define database connection for PostgreSQL
	fmt.Println(os.Getenv("DB_SERVER_URL"))
	// db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	db, err := sqlx.Connect("pgx", "host=john.db.elephantsql.com port=5432 user=tkaeaswb password=XHTYs5BSzLXESRRScK5fArMk2RNFJau- dbname=tkaeaswb sslmode=disable")
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
