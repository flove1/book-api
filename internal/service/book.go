package service

import (
	"context"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

func (m *Manager) CreateBook(ctx context.Context, book *entity.Book) error {
	err := m.Repository.CreateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetBookByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	book, err := m.Repository.GetBookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (m *Manager) GetBooks(ctx context.Context, title *string, author *string, tags *[]string, filter util.Filter) ([]*entity.Book, *util.Metadata, error) {
	filterSafeList := []string{
		"title",
		"author",
		"description",
		"rating",
		"created_at",
		"updated_at"}

	if !filter.ValidateSort(filterSafeList) {
		return nil, nil, ErrInvalidSortValue
	}

	books, meta, err := m.Repository.GetBooks(ctx, title, author, tags, filter)
	if err != nil {
		return nil, nil, err
	}

	return books, meta, nil
}

func (m *Manager) DeleteBook(ctx context.Context, articleID int64) error {
	err := m.Repository.DeleteBook(ctx, articleID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateBook(ctx context.Context, book *entity.Book) error {
	err := m.Repository.UpdateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) RefreshBooksRating(ctx context.Context) error {
	err := m.Repository.RefreshBooksRating(ctx)
	if err != nil {
		return err
	}

	return nil
}
