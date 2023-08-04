package pgrepo

import (
	"context"
	"errors"
	"fmt"
	"one-lab-final/internal/entity"
	"time"
)

func (p *Postgres) NewSuspension(ctx context.Context, suspension *entity.Suspension) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			reason,
			user_id,
			moderator_id,
			expires_in
		)
		VALUES ($1, $2, $3, $4)
	`, suspensionsTable)

	_, err := p.Pool.Exec(ctx, query, suspension.Reason, suspension.UserID, suspension.ModeratorID, suspension.ExpiresIn)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CheckSuspension(ctx context.Context, userID int64) ([]*entity.Suspension, error) {
	query := fmt.Sprintf(`
		SELECT 
			reason,
			moderator_id,
			expires_in,
			created_at,
			updated_at
		FROM %[1]s 
		WHERE 
			user_id = $1 
		AND 
			(created_at + expires_in) > $2
	`, suspensionsTable)

	suspensions := make([]*entity.Suspension, 0)

	rows, err := p.Pool.Query(ctx, query, userID, time.Now())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var suspension entity.Suspension
		err = rows.Scan(
			&suspension.Reason,
			&suspension.ModeratorID,
			&suspension.ExpiresIn,
			&suspension.CreatedAt,
			&suspension.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		suspensions = append(suspensions, &suspension)
	}

	return suspensions, nil
}

func (p *Postgres) UpdateSuspension(ctx context.Context, suspension *entity.Suspension) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			reason = COALESCE($1, reason),
			expires_in = COALESCE($2, expires_in),
			moderator_id = COALESCE($3, moderator_id),
			updated_at = $4
		WHERE 
			(expires_in + created_at) > $4
		AND
			id = $5
		AND 
			active
		
	`, suspensionsTable)

	tag, err := p.Pool.Exec(ctx, query,
		suspension.Reason,
		suspension.ExpiresIn,
		suspension.ModeratorID,
		time.Now(),
		suspension.ID,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("suspension either expired or does not exists")
	}

	return nil
}
