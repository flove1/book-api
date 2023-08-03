package repository

import (
	"context"

	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

//go:generate mockery --name DB
type Repository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByCredentials(ctx context.Context, username string) (*entity.User, error)
	GetUserByToken(ctx context.Context, token string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID int64) error

	CreateToken(ctx context.Context, token *entity.Token) error
	DeleteExpiredTokens(ctx context.Context) error

	CreateBook(ctx context.Context, book *entity.Book) error
	GetBookByID(ctx context.Context, bookID int64) (*entity.Book, error)
	GetBooks(ctx context.Context, title *string, author *string, tags *[]string, filter util.Filter) ([]*entity.Book, *util.Metadata, error)
	UpdateBook(ctx context.Context, book *entity.Book) error
	DeleteBook(ctx context.Context, articleID int64) error

	RefreshBooksRating(ctx context.Context) error

	CreateReview(ctx context.Context, review *entity.Review) error
	GetReviewsByBookID(ctx context.Context, bookID int64, filter util.Filter) ([]*entity.Review, *util.Metadata, error)
	UpdateReview(ctx context.Context, review *entity.Review) error
	DeleteReview(ctx context.Context, reviewID int64, userID int64) error
}
