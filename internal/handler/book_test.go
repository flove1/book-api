package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/service/mocks"
	"one-lab-final/pkg/util"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBook(t *testing.T) {
	var userID int64 = 123
	tests := []struct {
		Name         string
		RequestJSON  string
		MockResult   any
		ExpectedBook entity.Book
		ExpectedCode int
	}{
		{
			Name: "Create book successfully",
			RequestJSON: `
			{
				"title": "Book title",
				"author": "Great author",
				"description": "Boring description",
				"tags": ["tag1", "tag2", "tag3"],
				"year": 1994
			}`,
			MockResult: nil,
			ExpectedBook: entity.Book{
				Title:       util.StringToPointer("Book title"),
				Author:      util.StringToPointer("Great author"),
				Description: util.StringToPointer("Boring description"),
				Tags:        &[]string{"tag1", "tag2", "tag3"},
				Year:        1994,
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Name: "Empty tags",
			RequestJSON: `
			{
				"title": "Book title",
				"author": "Great author",
				"description": "Boring description",
				"tags": [""],
				"year": 1994
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Non-valid JSON",
			RequestJSON: `
			{
				"tags": ["tag1", "tag2", "tag3"]
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Error while saving entity",
			RequestJSON: `
			{
				"title": "Book title",
				"author": "Great author",
				"description": "Boring description",
				"tags": ["tag1", "tag2", "tag3"],
				"year": 1994
			}`,
			MockResult: errors.New("critical error"),
			ExpectedBook: entity.Book{
				Title:       util.StringToPointer("Book title"),
				Author:      util.StringToPointer("Great author"),
				Description: util.StringToPointer("Boring description"),
				Tags:        &[]string{"tag1", "tag2", "tag3"},
				Year:        1994,
			},
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("POST", "/books/new", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req
			ctx.Set("userID", userID)

			service.On("CreateBook", ctx, &test.ExpectedBook).Return(test.MockResult)
			handler.createBook(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestGetBookByID(t *testing.T) {
	var bookID int64 = 123
	tests := []struct {
		Name          string
		RequestURI    string
		MockResult    any
		MockResultErr error
		ExpectedID    int64
		ExpectedCode  int
	}{
		{
			Name:       "Retrieve book successfully",
			RequestURI: fmt.Sprintf("%d", bookID),
			MockResult: &entity.Book{
				Title:       util.StringToPointer("Title"),
				Description: util.StringToPointer("Description"),
				Author:      util.StringToPointer("Author"),
				Tags:        &[]string{"tag1", "tag2"},
				Rating:      100,
			},
			MockResultErr: nil,
			ExpectedID:    bookID,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:          "Non-valid URI",
			RequestURI:    "qwerty",
			MockResult:    nil,
			MockResultErr: errors.New("critical error"),
			ExpectedCode:  http.StatusBadRequest,
		},
		{
			Name:          "Error while retrieving entity",
			RequestURI:    fmt.Sprintf("%d", bookID),
			MockResult:    nil,
			MockResultErr: errors.New("critical error"),
			ExpectedID:    bookID,
			ExpectedCode:  http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("GET", "/books/"+test.RequestURI, strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("GetBookByID", ctx, test.ExpectedID).Return(test.MockResult, test.MockResultErr)
			handler.getBookByID(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestGetBooks(t *testing.T) {
	tests := []struct {
		Name           string
		RequestQuery   string
		MockResult     any
		MockResultErr  error
		MockResultMeta *util.Metadata
		ExpectedTitle  *string
		ExpectedAuthor *string
		ExpectedTags   *[]string
		ExpectedCode   int
	}{
		{
			Name:         "Get books succesfuly",
			RequestQuery: `page=1&page_size=10&sort=created_at&tags=tag1,tag2`,
			MockResult: []*entity.Book{
				{
					Title:       util.StringToPointer("Title"),
					Description: util.StringToPointer("Description"),
					Author:      util.StringToPointer("Author"),
					Tags:        &[]string{"tag1", "tag2"},
					Rating:      100,
				},
			},
			MockResultMeta: &util.Metadata{},
			MockResultErr:  nil,
			ExpectedTags:   &[]string{"tag1", "tag2"},
			ExpectedCode:   http.StatusOK,
		},
		{
			Name:         "Non-valid filter",
			RequestQuery: `page=none,page_size=many,sort=idk`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:           "Error while retriving",
			RequestQuery:   `page=1&page_size=10&sort=created_at`,
			MockResult:     nil,
			MockResultMeta: nil,
			MockResultErr:  errors.New("critical error"),
			ExpectedCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("GET", "/books?"+test.RequestQuery, &strings.Reader{})
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req

			service.On("GetBooks", ctx, test.ExpectedTitle, test.ExpectedAuthor, test.ExpectedTags, mock.AnythingOfType("util.Filter")).Return(test.MockResult, test.MockResultMeta, test.MockResultErr)
			handler.getBooks(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	var bookID int64 = 123
	tests := []struct {
		Name         string
		RequestURI   string
		RequestJSON  string
		MockResult   any
		ExpectedBook entity.Book
		ExpectedCode int
	}{
		{
			Name:       "Update book fully",
			RequestURI: fmt.Sprintf("%d", bookID),
			RequestJSON: `
			{
				"title": "Book title",
				"author": "Great author",
				"description": "Boring description",
				"tags": ["tag1", "tag2", "tag3"]
			}`,
			MockResult: nil,
			ExpectedBook: entity.Book{
				ID:          bookID,
				Title:       util.StringToPointer("Book title"),
				Author:      util.StringToPointer("Great author"),
				Description: util.StringToPointer("Boring description"),
				Tags:        &[]string{"tag1", "tag2", "tag3"},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:       "Update book partialy",
			RequestURI: fmt.Sprintf("%d", bookID),
			RequestJSON: `
			{
				"title": "Book title"
			}`,
			MockResult: nil,
			ExpectedBook: entity.Book{
				ID:    bookID,
				Title: util.StringToPointer("Book title"),
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:       "Empty tags",
			RequestURI: fmt.Sprintf("%d", bookID),
			RequestJSON: `
			{
				"tags": [""]
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "Non-valid URI path",
			RequestURI:   "qwerty",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Non-valid json body",
			RequestURI: fmt.Sprintf("%d", bookID),
			RequestJSON: `
			{
				OOPS
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Error while saving entity",
			RequestURI: fmt.Sprintf("%d", bookID),
			RequestJSON: `
			{
				"title": "Book title",
				"author": "Great author",
				"description": "Boring description",
				"tags": ["tag1", "tag2", "tag3"]
			}`,
			MockResult: errors.New("critical error"),
			ExpectedBook: entity.Book{
				ID:          bookID,
				Title:       util.StringToPointer("Book title"),
				Author:      util.StringToPointer("Great author"),
				Description: util.StringToPointer("Boring description"),
				Tags:        &[]string{"tag1", "tag2", "tag3"},
			},
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("PATCH", "/books/update/"+test.RequestURI, strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("UpdateBook", ctx, &test.ExpectedBook).Return(test.MockResult)
			handler.updateBook(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	var bookID int64 = 123
	tests := []struct {
		Name         string
		RequestURI   string
		MockResult   any
		ExpectedID   int64
		ExpectedCode int
	}{
		{
			Name:         "Delete book successfully",
			RequestURI:   fmt.Sprintf("%d", bookID),
			MockResult:   nil,
			ExpectedID:   bookID,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid URI path",
			RequestURI:   "qwerty",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "Error while deleting entity",
			RequestURI:   fmt.Sprintf("%d", bookID),
			MockResult:   errors.New("critical error"),
			ExpectedID:   bookID,
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("DELETE", "/reviews/delete/"+test.RequestURI, strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("DeleteBook", ctx, test.ExpectedID).Return(test.MockResult)
			handler.deleteBook(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
