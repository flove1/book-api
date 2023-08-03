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

func TestCreateReview(t *testing.T) {
	var userID int64 = 123
	var bookID int64 = 123
	tests := []struct {
		Name           string
		RequestJSON    string
		MockResult     any
		ExpectedReview entity.Review
		ExpectedCode   int
	}{
		{
			Name: "Create review successfully",
			RequestJSON: `
			{
				"content": "This is a review", 
				"rating": 5, 
				"book_id": 123
			}`,
			MockResult: nil,
			ExpectedReview: entity.Review{
				Content: util.StringToPointer("This is a review"),
				Rating:  util.IntToPointer(5),
				BookID:  bookID,
				UserID:  userID,
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Name: "Create review with invalid JSON",
			RequestJSON: `
			{
				"not_content": "This is a review", 
				"not_rating": 5, 
				"not_book_id": 123
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Error while saving entity",
			RequestJSON: `
			{
				"content": "This is a review", 
				"rating": 5, 
				"book_id": 123
			}`,
			MockResult: errors.New("critical error"),
			ExpectedReview: entity.Review{
				Content: util.StringToPointer("This is a review"),
				Rating:  util.IntToPointer(5),
				BookID:  123,
				UserID:  userID,
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

			req, _ := http.NewRequest("POST", "/reviews/new", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req
			ctx.Set("userID", userID)

			service.On("CreateReview", ctx, &test.ExpectedReview).Return(test.MockResult)
			handler.createReview(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestGetReviewsByBookID(t *testing.T) {
	var bookID int64 = 123
	tests := []struct {
		Name           string
		RequestURI     string
		RequestQuery   string
		MockResultErr  error
		MockResultMeta *util.Metadata
		MockResult     any
		ExpectedID     int64
		ExpectedCode   int
	}{
		{
			Name:         "Get reviews succesfuly",
			RequestURI:   fmt.Sprintf("%d", bookID),
			RequestQuery: `page=1&page_size=10&sort=created_at`,
			MockResult: []*entity.Review{
				{
					ID:      1,
					Content: util.StringToPointer("Title 1"),
					Rating:  util.IntToPointer(100),
					UserID:  1,
					BookID:  bookID,
				},
				{
					ID:      2,
					Content: util.StringToPointer("Title 2"),
					Rating:  util.IntToPointer(25),
					UserID:  2,
					BookID:  bookID,
				},
			},
			MockResultMeta: &util.Metadata{},
			MockResultErr:  nil,
			ExpectedID:     bookID,
			ExpectedCode:   http.StatusOK,
		},
		{
			Name:         "Non-valid URI path",
			RequestURI:   "qwerty",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "Non-valid filter",
			RequestURI:   fmt.Sprintf("%d", bookID),
			RequestQuery: `page=none,page_size=many,sort=idk`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:           "Error while retriving",
			RequestURI:     fmt.Sprintf("%d", bookID),
			RequestQuery:   `page=1&page_size=10&sort=created_at`,
			MockResult:     nil,
			MockResultMeta: nil,
			MockResultErr:  errors.New("critical error"),
			ExpectedID:     bookID,
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

			req, _ := http.NewRequest("GET", "/books/reviews/"+test.RequestURI+"/reviews?"+test.RequestQuery, &strings.Reader{})
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("GetReviewsByBookID", ctx, test.ExpectedID, mock.AnythingOfType("util.Filter")).Return(test.MockResult, test.MockResultMeta, test.MockResultErr)
			handler.getReviewsByBookID(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestUpdateReview(t *testing.T) {
	var userID int64 = 123
	var reviewID int64 = 123
	tests := []struct {
		Name           string
		RequestURI     string
		RequestJSON    string
		MockResult     any
		ExpectedReview entity.Review
		ExpectedCode   int
	}{
		{
			Name:       "Update review fully",
			RequestURI: fmt.Sprintf("%d", reviewID),
			RequestJSON: `
			{
				"content": "This is a review", 
				"rating": 5
			}`,
			MockResult: nil,
			ExpectedReview: entity.Review{
				ID:      reviewID,
				Content: util.StringToPointer("This is a review"),
				Rating:  util.IntToPointer(5),
				UserID:  userID,
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:       "Update review partialy",
			RequestURI: fmt.Sprintf("%d", reviewID),
			RequestJSON: `
			{
				"rating": 5
			}`,
			MockResult: nil,
			ExpectedReview: entity.Review{
				ID:     reviewID,
				Rating: util.IntToPointer(5),
				UserID: userID,
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid URI path",
			RequestURI:   "qwerty",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Non-valid json body",
			RequestURI: fmt.Sprintf("%d", reviewID),
			RequestJSON: `
			{
				OOPS
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Error while saving entity",
			RequestURI: fmt.Sprintf("%d", reviewID),
			RequestJSON: `
			{
				"content": "This is a review", 
				"rating": 5
			}`,
			MockResult: errors.New("critical error"),
			ExpectedReview: entity.Review{
				ID:      reviewID,
				Content: util.StringToPointer("This is a review"),
				Rating:  util.IntToPointer(5),
				UserID:  userID,
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

			req, _ := http.NewRequest("PATCH", "/reviews/update/"+test.RequestURI, strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Set("userID", userID)
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("UpdateReview", ctx, &test.ExpectedReview).Return(test.MockResult)
			handler.updateReview(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestDeleteReview(t *testing.T) {
	var userID int64 = 123
	var reviewID int64 = 123
	tests := []struct {
		Name         string
		RequestURI   string
		MockResult   any
		ExpectedID   int64
		ExpectedCode int
	}{
		{
			Name:         "Delete review succesfuly",
			RequestURI:   fmt.Sprintf("%d", reviewID),
			MockResult:   nil,
			ExpectedID:   reviewID,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid URI path",
			RequestURI:   "qwerty",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "Error while deleting entity",
			RequestURI:   fmt.Sprintf("%d", reviewID),
			MockResult:   errors.New("critical error"),
			ExpectedID:   reviewID,
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
			ctx.Set("userID", userID)
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("DeleteReview", ctx, test.ExpectedID, userID).Return(test.MockResult)
			handler.deleteReview(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
