package pgrepo

import (
	"one-lab-final/internal/config"

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
	Pool   *pgxpool.Pool
	Config *config.Config
}

func New(pool *pgxpool.Pool, config *config.Config) *Postgres {
	return &Postgres{
		Pool:   pool,
		Config: config,
	}
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
