CREATE MATERIALIZED VIEW books_avg_rating_view AS SELECT AVG(rating) AS rating, book_id FROM reviews GROUP BY book_id;

CREATE INDEX idx_books_avg_rating_view ON books_avg_rating_view (rating);