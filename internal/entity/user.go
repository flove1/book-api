package entity

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  *string   `json:"username" db:"username"`
	Email     *string   `json:"-" db:"email"`
	FirstName *string   `json:"first_name" db:"first_name"`
	LastName  *string   `json:"last_name" db:"last_name"`
	Password  Password  `json:"-" db:"password_hash"`
	Role      Role      `json:"role" db:"role"`
	Suspended bool      `json:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Password struct {
	Plaintext *string `json:"-"`
	Hash      *[]byte `json:"-" db:"password_hash"`
}
