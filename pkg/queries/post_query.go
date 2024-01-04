package queries

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

type PostQueries struct {
	*sqlx.DB
}

// GetPosts method for getting all posts.
func (q *PostQueries) GetPosts(id uuid.UUID) ([]models.Post, error) {
	// Define posts variable.
	posts := []models.Post{}

	// Define query string.
	query := `SELECT * FROM posts WHERE user_id = $1`

	// Send query to database.
	err := q.Select(&posts, query, id)
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

// CreatePost method for creating post by given Post object.
func (q *PostQueries) CreatePost(u *models.Post) error {
	// Define query string.
	query := `INSERT INTO posts (id, user_id, title, details, is_publish, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(query, u.ID, u.UserId, u.Title, u.Details, u.IsPublish, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// CreatePost method for creating post by given Post object.
func (q *PostQueries) UpdatePost(input *models.PostAndUserId, u *models.PostUpdate) error {
	// Define base query string.
	query := `UPDATE posts SET title = $3`
	// Initialize parameters slice with the name value.
	params := []interface{}{input.ID, input.UserId, u.Title}

	// Conditionally include gender in the query if it's provided.
	if u.Details != "" {
		query += `, details = $4`
		// Append gender value to the parameters slice.
		params = append(params, u.Details)
	}
	fmt.Println(u.IsPublish)
	if u.IsPublish {
		query += `, is_publish = $5`
		// Append gender value to the parameters slice.
		params = append(params, u.IsPublish)
	} else {
		query += `, is_publish = $5`
		// Append gender value to the parameters slice.
		params = append(params, u.IsPublish)
	}

	// Include the common part of the query.
	query += `, updated_at = $` + strconv.Itoa(len(params)+1) + ` WHERE id = $1 AND user_id = $2`

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
func (q *PostQueries) DeletePost(id uuid.UUID, userId uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM posts WHERE id = $1 AND user_id = $2`

	// Send query to database.
	_, err := q.Exec(query, id, userId)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
