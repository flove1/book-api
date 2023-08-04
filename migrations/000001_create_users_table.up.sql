CREATE EXTENSION IF NOT EXISTS citext;

CREATE TYPE role AS ENUM('USER', 'MODERATOR', 'ADMIN');

CREATE TABLE IF NOT EXISTS users (
	id bigserial PRIMARY KEY,
	username text UNIQUE NOT NULL,
	email citext UNIQUE NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	password_hash bytea NOT NULL,
    role role NOT NULL DEFAULT 'USER',
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);