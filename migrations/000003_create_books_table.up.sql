CREATE TABLE IF NOT EXISTS books (
	id bigserial PRIMARY KEY,
	title text NOT NULL,
	author text NOT NULL,
    description text NOT NULL,
    tags citext[] NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);