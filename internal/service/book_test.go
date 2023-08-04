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
	tests := []struct {
		Name         string
		MockResult   any
		ExpectedBook *entity.Book
	}{
		{
			Name:         "Book created successfully",
			MockResult:   nil,
			ExpectedBook: &entity.Book{},
		},
		{
			Name:         "Some error ocurred while saving",
			MockResult:   errors.New("critical error"),
			ExpectedBook: &entity.Book{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("CreateBook", ctx, test.ExpectedBook).Return(test.MockResult)

			assert.Equal(t, service.CreateBook(ctx, test.ExpectedBook), test.MockResult)
		})
	}
}

func TestGetBookByID(t *testing.T) {
	var bookID int64 = 123
	tests := []struct {
		Name        string
		MockResult  any
		MockError   error
		ExpectError bool
	}{
		{
			Name:        "Book retrieved successfully",
			MockResult:  nil,
			ExpectError: false,
		},
		{
			Name:        "Some error ocurred while retriving",
			MockError:   errors.New("critical error"),
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("GetBookByID", ctx, bookID).Return(test.MockResult, test.MockError)

			_, err := service.GetBookByID(ctx, bookID)

			if test.ExpectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
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
	title := "Title"
	author := "Author"
	tags := []string{"tag1"}
	tests := []struct {
		Name           string
		MockResultErr  error
		MockResultMeta *util.Metadata
		MockResult     any
		SkipMock       bool
		BookID         int64
		Filter         util.Filter
		ExpectedErr    bool
	}{
		{
			Name:           "Get books succesfuly",
			MockResult:     []*entity.Book{{}, {}},
			MockResultMeta: &util.Metadata{},
			Filter:         util.NewFilter(1, 10, "created_at"),
			ExpectedErr:    false,
		},
		{
			Name:        "Non-valid sort",
			SkipMock:    true,
			Filter:      util.NewFilter(1, 10, "everything"),
			ExpectedErr: true,
		},
		{
			Name:          "Error while retriving",
			MockResultErr: errors.New("critical error"),
			Filter:        util.NewFilter(1, 10, "created_at"),
			ExpectedErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			if !test.SkipMock {
				repo.On("GetBooks", ctx, &title, &author, &tags, test.Filter).Return(test.MockResult, test.MockResultMeta, test.MockResultErr)
			}

			_, _, err := service.GetBooks(ctx, &title, &author, &tags, test.Filter)

			if test.ExpectedErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
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
	tests := []struct {
		Name         string
		MockResult   any
		ExpectedBook *entity.Book
	}{
		{
			Name:         "Book updated successfully",
			MockResult:   nil,
			ExpectedBook: &entity.Book{},
		},
		{
			Name:         "Some error ocurred while updating",
			MockResult:   errors.New("critical error"),
			ExpectedBook: &entity.Book{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("UpdateBook", ctx, test.ExpectedBook).Return(test.MockResult)

			assert.Equal(t, service.UpdateBook(ctx, test.ExpectedBook), test.MockResult)
		})
	}
}

func TestRefreshBooksRating(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
	}{
		{
			Name:       "Refreshed successfully",
			MockResult: nil,
		},
		{
			Name:       "Some error ocurred while refreshing",
			MockResult: errors.New("critical error"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("RefreshBooksRating", ctx).Return(test.MockResult)

			assert.Equal(t, service.RefreshBooksRating(ctx), test.MockResult)
		})
	}
}
