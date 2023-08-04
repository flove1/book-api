package entity

import "time"

type Suspension struct {
	ID          int64          `json:"id" db:"id"`
	Reason      *string        `json:"reason" db:"reason"`
	UserID      int64          `json:"user_id" db:"user_id"`
	ModeratorID int64          `json:"moderator_id" db:"moderator_id"`
	ExpiresIn   *time.Duration `json:"expires_in" db:"expires_in" swaggertype:"primitive,integer"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}
