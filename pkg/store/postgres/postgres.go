package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Settings struct {
	host     string
	username string
	password string
	port     string
	db       string
}

func (s *Settings) parseToDSN() string {
	q := url.Values{}
	q.Add("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(s.username, s.password),
		Host:     fmt.Sprintf("%s:%s", s.host, s.port),
		Path:     s.db,
		RawQuery: q.Encode(),
	}

	return u.String()
}

func ConnectDB(options ...Option) (*pgxpool.Pool, error) {
	s := new(Settings)

	for _, option := range options {
		option(s)
	}

	config, err := pgxpool.ParseConfig(s.parseToDSN())
	if err != nil {
		return nil, fmt.Errorf("pgxpool parse config err: %w", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect err: %w", err)
	}

	return conn, nil
}
