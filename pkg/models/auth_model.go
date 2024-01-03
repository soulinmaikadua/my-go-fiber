package models

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name      string    `db:"name" json:"name"`
	Gender    string    `db:"gender" json:"gender"`
	Email     string    `db:"email" json:"email" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Token     string    `json:"token"`
}
