package configs

import "github.com/soulinmaikadua/my-go-fiber/pkg/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries
	*queries.PostQueries
}


func OpenDBConnection() (*Queries, error) {
	// Define a new Postgres connection
	db, err := PostgresConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries for models
		UserQueries: &queries.UserQueries{DB: db},
		PostQueries: &queries.PostQueries{DB: db},
	}, nil
}