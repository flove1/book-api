package entity

import "time"

type Token struct {
	Plaintext string    `json:"token"`
	Expiry    time.Time `json:"expiry" db:"expiry"`
	UserID    int64     `json:"-" db:"user_id"`
	Hash      []byte    `json:"-" db:"token"`
}
