package entity

import "time"

type Review struct {
	ID        int64     `json:"id" db:"id"`
	Content   *string   `json:"content" db:"content"`
	Rating    *int64    `json:"rating" db:"rating"`
	UserID    int64     `json:"user_id" db:"user_id"`
	BookID    int64     `json:"book_id" db:"book_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
