package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type PostQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *PostQueries) GetPosts() ([]models.Post, error) {
	// Define books variable.
	posts := []models.Post{}

	// Define query string.
	query := `SELECT * FROM posts`

	// Send query to database.
	err := q.Select(&posts, query)
	if err != nil {
		// Return empty object and error.
		return posts, err
	}

	// Return query result.
	return posts, nil
}