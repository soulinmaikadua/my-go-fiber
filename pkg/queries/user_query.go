package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type UserQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *UserQueries) GetUsers() ([]models.User, error) {
	// Define books variable.
	books := []models.User{}

	// Define query string.
	query := `SELECT id, name, gender, email, created_at, updated_at FROM users`

	// Send query to database.
	err := q.Select(&books, query)
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetUser method for getting one user by given ID.
func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	// Define user variable.
	user := models.User{}
	// fmt.Println(id.String())

	// Define query string.
	query := `SELECT id, name, gender, email, created_at, updated_at FROM users WHERE id = $1`

	// fmt.Println("SQL Query:", query)

	// Send query to database.
	err := q.Get(&user, query, id)
	if err != nil {
		// fmt.Println("Error executing query:", err)
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser method for creating user by given User object.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users (id, name, gender, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(query, u.ID, u.Name, u.Gender, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
