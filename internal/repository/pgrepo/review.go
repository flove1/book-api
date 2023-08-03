package pgrepo

import (
	"context"
	"errors"
	"fmt"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
	"time"
)

func (p *Postgres) CreateReview(ctx context.Context, review *entity.Review) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			content,
			rating,
			user_id,
			book_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`, reviewsTable)

	err := p.Pool.QueryRow(ctx, query, review.Content, review.Rating, review.UserID, review.BookID).Scan(&review.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetReviewsByBookID(ctx context.Context, bookID int64, filter util.Filter) ([]*entity.Review, *util.Metadata, error) {
	totalQuery := fmt.Sprintf(`
	SELECT 
		count(*) AS total_count
	FROM %s
	WHERE book_id = $1	
	`, reviewsTable)

	dataQuery := fmt.Sprintf(`
	SELECT 
		id,
		content,
		rating,
		user_id,
		book_id,
		created_at,
		updated_at
	FROM %[1]s
	WHERE book_id = $1
	ORDER BY %[2]s %[3]s
	LIMIT $2 OFFSET $3
`, reviewsTable, filter.FormatSort(), filter.SortDirection())

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	defer tx.Rollback(ctx)

	var totalCount int
	reviews := make([]*entity.Review, 0)

	err = tx.QueryRow(ctx, totalQuery, bookID).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	rows, err := tx.Query(ctx, dataQuery, bookID, filter.Limit(), filter.Offset())
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var review entity.Review
		rows.Scan(
			&review.ID,
			&review.Content,
			&review.Rating,
			&review.UserID,
			&review.BookID,
			&review.CreatedAt,
			&review.UpdatedAt,
		)

		reviews = append(reviews, &review)
	}

	metadata := filter.CalculateMetadata(totalCount)

	return reviews, &metadata, nil
}

func (p *Postgres) UpdateReview(ctx context.Context, review *entity.Review) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			content = COALESCE($2, content),
			rating = COALESCE($3, rating),
			updated_at = $4
		WHERE 
			id = $1
	`, reviewsTable)

	tag, err := p.Pool.Exec(ctx, query, review.ID, review.Content, review.Rating, time.Now())
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("article does not exists or does not belong to user")
	}

	return nil
}

func (p *Postgres) DeleteReview(ctx context.Context, reviewID int64, userID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE 
			id = $1
		AND
			user_id = $2
	`, reviewsTable)

	_, err := p.Pool.Exec(ctx, query, reviewID, userID)
	if err != nil {
		return err
	}

	return nil
}
