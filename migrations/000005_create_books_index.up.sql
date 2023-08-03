CREATE INDEX IF NOT EXISTS idx_books_title ON books USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS idx_books_title ON books USING GIN (to_tsvector('simple', author));
CREATE INDEX IF NOT EXISTS idx_books_tags ON books USING GIN (tags);
