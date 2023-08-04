package pgrepo

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository"

	"github.com/jackc/pgx/v4"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			username,
			email,
			first_name,
			last_name,
			password_hash
			)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`, usersTable)

	err := p.Pool.QueryRow(ctx, query, u.Username, u.Email, u.FirstName, u.LastName, u.Password.Hash).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := fmt.Sprintf(`
			SELECT 
				id,
				username,
				email,
				first_name,
				last_name,
				created_at
			FROM %s
			WHERE 
				username = $1
			`, usersTable)

	user := &entity.User{}

	err := p.Pool.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, repository.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (p *Postgres) GetUserByCredentials(ctx context.Context, credentials string) (*entity.User, error) {
	query := fmt.Sprintf(`
			SELECT 
				id,
				username,
				email,
				first_name,
				last_name,
				password_hash
			FROM %s
			WHERE 
				username = $1
			OR
				email = $1
			`, usersTable)

	user := &entity.User{}

	err := p.Pool.QueryRow(ctx, query, credentials).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password.Hash,
	)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, repository.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (p *Postgres) GetUserByToken(ctx context.Context, token string) (*entity.User, error) {
	hash := sha256.Sum256([]byte(token))

	query := fmt.Sprintf(`
		SELECT 
			u.id,
			u.username,
			u.email,
			u.first_name,
			u.last_name,
			u.password_hash,
			u.created_at,
			u.updated_at,
			u.role
		FROM %[1]s u
		INNER JOIN %[2]s  t
		ON u.id = t.user_id
		WHERE t.hash = $1
		AND t.expiry > $2
	`, usersTable, tokensTable, suspensionsTable)

	user := new(entity.User)
	var roleString string

	err := p.Pool.QueryRow(ctx, query, hash[:], time.Now()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&roleString,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, repository.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	role, ok := entity.StringToRole(roleString)
	if !ok {
		return nil, errors.New("error while parsing role")
	}

	user.Role = role

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			first_name = COALESCE($2, first_name),
			last_name = COALESCE($3, last_name),
			password_hash = COALESCE($4, password_hash)
		WHERE 
			id = $1
	`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.ID, u.FirstName, u.LastName, u.Password.Hash)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, userID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE 
			id = $1
	`, usersTable)

	_, err := p.Pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
