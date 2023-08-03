package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable         = "users"
	booksTable         = "books"
	tokensTable        = "tokens"
	reviewsTable       = "reviews"
	suspensionsTable   = "suspensions"
	booksAvgRatingView = "books_avg_rating_view"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{pool}
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
