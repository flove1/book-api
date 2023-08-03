package pgrepo

import (
	"context"
	"fmt"
	"one-lab-final/internal/entity"
	"time"
)

func (p *Postgres) CreateToken(ctx context.Context, token *entity.Token) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			hash,
			user_id,
			expiry
		)
		VALUES ($1, $2, $3)
	`, tokensTable)

	_, err := p.Pool.Exec(ctx, query, token.Hash, token.UserID, token.Expiry)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteExpiredTokens(ctx context.Context) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE expiry < $1
	`, tokensTable)

	_, err := p.Pool.Exec(ctx, query, time.Now())
	if err != nil {
		return err
	}
	return nil
}
