package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type AuthQueries struct {
	*sqlx.DB
}

// CreateUser method for creating user by given User object.
func (q *AuthQueries) SignUp(u *models.User) error {

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

// SignInUser method for creating user by given User object.
func (q *AuthQueries) SignIn(u *models.User) error {

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	_, err := q.Query(query, u.Email)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
