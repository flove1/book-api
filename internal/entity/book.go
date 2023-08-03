package entity

import "time"

type Book struct {
	ID          int64     `json:"id" db:"id"`
	Title       *string   `json:"title" db:"title"`
	Description *string   `json:"description" db:"description"`
	Author      *string   `json:"author" db:"author"`
	Tags        *[]string `json:"tags" db:"tags"`
	Rating      float32   `json:"rating"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
