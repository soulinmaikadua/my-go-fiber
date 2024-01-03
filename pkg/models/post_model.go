package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	UserId    string    `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title     string    `db:"title" json:"title" validate:"required"`
	Body      string    `db:"body" json:"body" validate:"required"`
	Publish   bool      `db:"publish" json:"publish"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
