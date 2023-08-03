CREATE TABLE IF NOT EXISTS reviews (
	id bigserial PRIMARY KEY,
    content text NOT NULL,
    rating integer DEFAULT 0,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    book_id bigint NOT NULL REFERENCES books ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, book_id)
)