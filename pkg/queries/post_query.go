package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type PostQueries struct {
	*sqlx.DB
}

// GetPosts method for getting all posts.
func (q *PostQueries) GetPosts() ([]models.Post, error) {
	// Define posts variable.
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

// GetPost method for getting one post by given ID.
func (q *PostQueries) GetPost(id uuid.UUID) (models.Post, error) {
	// Define user variable.
	post := models.Post{}
	// fmt.Println(id.String())

	// Define query string.
	query := `SELECT * FROM posts WHERE id = $1`

	// fmt.Println("SQL Query:", query)

	// Send query to database.
	err := q.Get(&post, query, id)
	if err != nil {
		// fmt.Println("Error executing query:", err)
		// Return empty object and error.
		return post, err
	}

	// Return query result.
	return post, nil
}
