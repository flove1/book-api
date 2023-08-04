package service

import (
	"context"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

func (m *Manager) CreateBook(ctx context.Context, book *entity.Book) error {
	return m.Repository.CreateBook(ctx, book)
}

func (m *Manager) GetBookByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	return m.Repository.GetBookByID(ctx, bookID)
}

func (m *Manager) GetBooks(ctx context.Context, title *string, author *string, tags *[]string, filter util.Filter) ([]*entity.Book, *util.Metadata, error) {
	filterSafeList := []string{
		"title",
		"author",
		"description",
		"rating",
		"year",
		"created_at",
		"updated_at",
	}

	if !filter.ValidateSort(filterSafeList) {
		return nil, nil, ErrInvalidSortValue
	}

	return m.Repository.GetBooks(ctx, title, author, tags, filter)
}

func (m *Manager) DeleteBook(ctx context.Context, articleID int64) error {
	return m.Repository.DeleteBook(ctx, articleID)
}

func (m *Manager) UpdateBook(ctx context.Context, book *entity.Book) error {
	return m.Repository.UpdateBook(ctx, book)
}

func (m *Manager) RefreshBooksRating(ctx context.Context) error {
	return m.Repository.RefreshBooksRating(ctx)
}
