package service

import (
	"context"
	"errors"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository/mocks"
	"one-lab-final/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	title := "Book"
	book := new(entity.Book)
	book.Title = &title

	repo.On("CreateBook", ctx, book).Return(nil).Once()
	repo.On("CreateBook", ctx, book).Return(errors.New("")).Once()

	result := service.CreateBook(ctx, book)
	assert.Nil(t, result)

	result = service.CreateBook(ctx, book)
	assert.NotNil(t, result)
}

func TestGetBookByID(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	book := new(entity.Book)

	repo.On("GetBookByID", ctx, int64(5)).Return(book, nil)
	repo.On("GetBookByID", ctx, int64(42)).Return(nil, errors.New("some error"))

	result, err := service.GetBookByID(ctx, 5)
	assert.Equal(t, book, result)
	assert.Nil(t, err)

	result, err = service.GetBookByID(ctx, 42)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestGetBooks(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	books := make([]*entity.Book, 5)

	filter := util.NewFilter(1, 50, "created_at")
	invalidFilter := util.NewFilter(1, 50, "name")

	title := "Book"
	author := "Author"
	tags := []string{"tag1", "tag2"}

	meta := filter.CalculateMetadata(len(books))

	repo.On("GetBooks", ctx, &title, &author, &tags, filter).Return(books, &meta, nil).Once()
	repo.On("GetBooks", ctx, &title, &author, &tags, filter).Return(nil, nil, errors.New("some error")).Once()

	result, _, err := service.GetBooks(ctx, &title, &author, &tags, filter)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	result, _, err = service.GetBooks(ctx, &title, &author, &tags, invalidFilter)
	assert.Nil(t, result)
	assert.NotNil(t, err)

	result, _, err = service.GetBooks(ctx, &title, &author, &tags, filter)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestDeleteBook(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	var id int64 = 5

	repo.On("DeleteBook", ctx, id).Return(nil).Once()
	repo.On("DeleteBook", ctx, id).Return(errors.New("some error")).Once()

	result := service.DeleteBook(ctx, id)
	assert.Nil(t, result)

	result = service.DeleteBook(ctx, id)
	assert.NotNil(t, result)
}

func TestUpdateBook(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	title := "Book"
	book := new(entity.Book)
	book.ID = 1
	book.Title = &title

	repo.On("UpdateBook", ctx, book).Return(nil).Once()
	repo.On("UpdateBook", ctx, book).Return(errors.New("some error")).Once()

	result := service.UpdateBook(ctx, book)
	assert.Nil(t, result)

	result = service.UpdateBook(ctx, book)
	assert.NotNil(t, result)
}

func TestRefreshBooksRating(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := New(repo, nil)
	ctx := context.Background()

	repo.On("RefreshBooksRating", ctx).Return(nil).Once()
	repo.On("RefreshBooksRating", ctx).Return(errors.New("some error")).Once()

	result := service.RefreshBooksRating(ctx)
	assert.Nil(t, result)

	result = service.RefreshBooksRating(ctx)
	assert.NotNil(t, result)
}
