package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name   string    `db:"name" json:"name"`
	Gender string    `db:"gender" json:"gender"`
	// FirstName string `db:"first_name" json:"first_name" validate:"required"`
	// LastName string `db:"last_name" json:"last_name" validate:"required"`
	// UserName string `db:"username" json:"username" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required"`
	Password  string    `db:"password" json:"password" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
