package queries

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type UserQueries struct {
	*sqlx.DB
}

// GetUsers method for getting all users.
func (q *UserQueries) GetUsers() ([]models.User, error) {
	// Define users variable.
	users := []models.User{}

	// Define query string.
	query := `SELECT id, name, gender, email, created_at, updated_at FROM users`

	// Send query to database.
	err := q.Select(&users, query)
	if err != nil {
		// Return empty object and error.
		return users, err
	}

	// Return query result.
	return users, nil
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

// UpdateUser method for updating user by given User object.
func (q *UserQueries) UpdateUser(id uuid.UUID, u *models.UpdateUser) error {
	// Define base query string.
	query := `UPDATE users SET name = $2`
	// Initialize parameters slice with the name value.
	params := []interface{}{id, u.Name}

	// Conditionally include gender in the query if it's provided.
	if u.Gender != "" {
		query += `, gender = $3`
		// Append gender value to the parameters slice.
		params = append(params, u.Gender)
	}

	// Include the common part of the query.
	query += `, updated_at = $` + strconv.Itoa(len(params)+1) + ` WHERE id = $1`

	// Append the updated_at value to the parameters slice.
	params = append(params, u.UpdatedAt)

	// Send query to the database.
	_, err := q.Exec(query, params...)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteUser method for delete user by given ID.
func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM users WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
