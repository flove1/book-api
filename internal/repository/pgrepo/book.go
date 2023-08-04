package pgrepo

import (
	"context"
	"errors"
	"fmt"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository"
	"one-lab-final/pkg/util"
	"strings"
	"time"
)

func (p *Postgres) CreateBook(ctx context.Context, book *entity.Book) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			title,
			author,
			description,
			tags,
			year
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`, booksTable)

	err := p.Pool.QueryRow(ctx, query, book.Title, book.Author, book.Description, book.Tags, book.Year).Scan(&book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetBookByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	query := fmt.Sprintf(`
		SELECT 
			b.id,
			b.title,
			b.description,
			b.author,
			b.tags,
			b.year,
			b.created_at,
			b.updated_at,
			COALESCE((SELECT rating FROM %[2]s r WHERE r.book_id = b.id), 0) AS rating
		FROM %[1]s b
		WHERE 
			b.id = $1
	`, booksTable, booksAvgRatingView)

	book := entity.Book{}

	tags := ""
	err := p.Pool.QueryRow(ctx, query, bookID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.Author,
		&tags,
		&book.Year,
		&book.CreatedAt,
		&book.UpdatedAt,

		&book.Rating,
	)
	if err != nil {
		return nil, err
	}

	tagsArray := strings.Split(tags[1:len(tags)-1], ",")
	book.Tags = &tagsArray

	return &book, nil
}

func (p *Postgres) GetBooks(ctx context.Context, title *string, author *string, tags *[]string, filter util.Filter) ([]*entity.Book, *util.Metadata, error) {
	totalQuery := fmt.Sprintf(`
	SELECT 
		count(*) AS total_count
	FROM %s
	WHERE 
		(to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 IS NULL)
	AND 
		(tags @> $2 OR $2 IS NULL)
	
	`, booksTable)

	dataQuery := fmt.Sprintf(`
		SELECT 
			b.id,
			b.title,
			b.description,
			b.author,
			b.tags,
			b.year,
			b.created_at,
			b.updated_at,
			COALESCE((SELECT avg(rating) FROM %[2]s r WHERE r.book_id = b.id), 0) AS rating
		FROM %[1]s b
		WHERE 
			(to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 IS NULL)
		AND
			(to_tsvector('simple', author) @@ plainto_tsquery('simple', $2) OR $2 IS NULL)
		AND 
			(tags @> $3 OR $3 IS NULL)
		ORDER BY %[3]s %[4]s, id ASC
		LIMIT $4 OFFSET $5
	`, booksTable, booksAvgRatingView, filter.FormatSort(), filter.SortDirection())

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	defer tx.Rollback(ctx)

	var totalCount int
	books := make([]*entity.Book, 0)

	err = tx.QueryRow(ctx, totalQuery, title, tags).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	rows, err := tx.Query(ctx, dataQuery, title, author, tags, filter.Limit(), filter.Offset())
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book entity.Book
		tags := ""
		rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Author,
			&tags,
			&book.Year,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Rating,
		)
		tagsArray := strings.Split(tags[1:len(tags)-1], ",")
		book.Tags = &tagsArray

		books = append(books, &book)
	}

	metadata := filter.CalculateMetadata(totalCount)

	return books, &metadata, nil
}

func (p *Postgres) DeleteBook(ctx context.Context, articleID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, booksTable)

	tag, err := p.Pool.Exec(ctx, query, articleID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateBook(ctx context.Context, book *entity.Book) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			title = COALESCE($1, title),
			author = COALESCE($2, author),
			description = COALESCE($3, description),
			tags = COALESCE($4, tags),
			updated_at = $5
		WHERE 
			id = $6
	`, booksTable)

	tag, err := p.Pool.Exec(ctx, query,
		book.Title,
		book.Author,
		book.Description,
		book.Tags,
		time.Now(),
		book.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("book doesn't exist")
	}

	return nil
}

func (p *Postgres) RefreshBooksRating(ctx context.Context) error {
	query := fmt.Sprintf(`
		REFRESH MATERIALIZED VIEW %s;
	`, booksAvgRatingView)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
