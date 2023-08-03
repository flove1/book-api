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

func TestCreateReview(t *testing.T) {
	var userID int64 = 5
	var bookID int64 = 42
	tests := []struct {
		Name           string
		MockResult     any
		ExpectedReview *entity.Review
	}{
		{
			Name:       "Review created successfullt",
			MockResult: nil,
			ExpectedReview: &entity.Review{
				Content: util.StringToPointer("Very good"),
				Rating:  util.IntToPointer(100),
				UserID:  userID,
				BookID:  bookID,
			},
		},
		{
			Name:       "Some error ocurred while saving",
			MockResult: errors.New("critical error"),
			ExpectedReview: &entity.Review{
				Content: util.StringToPointer("Very good"),
				Rating:  util.IntToPointer(100),
				UserID:  userID,
				BookID:  bookID,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("CreateReview", ctx, test.ExpectedReview).Return(test.MockResult)

			assert.Equal(t, service.CreateReview(ctx, test.ExpectedReview), test.MockResult)
		})
	}
}

// func TestGetReviewsForBookID(t *testing.T) {
// 	tests := []struct {
// 		Name           string
// 		RequestURI     string
// 		RequestQuery   string
// 		MockResultErr  error
// 		MockResultMeta *util.Metadata
// 		MockResult     any
// 		Filter         util.Filter
// 		ExpectedCode   int
// 	}{
// 		{
// 			Name:         "Get reviews succesfuly",
// 			RequestURI:   fmt.Sprintf("%d", bookID),
// 			RequestQuery: `page=1&page_size=10&sort=created_at`,
// 			MockResult: []*entity.Review{
// 				{
// 					ID:      1,
// 					Content: util.StringToPointer("Title 1"),
// 					Rating:  util.IntToPointer(100),
// 					UserID:  1,
// 					BookID:  bookID,
// 				},
// 				{
// 					ID:      2,
// 					Content: util.StringToPointer("Title 2"),
// 					Rating:  util.IntToPointer(25),
// 					UserID:  2,
// 					BookID:  bookID,
// 				},
// 			},
// 			MockResultMeta: &util.Metadata{},
// 			MockResultErr:  nil,
// 			ExpectedID:     bookID,
// 			ExpectedCode:   http.StatusOK,
// 		},
// 		{
// 			Name:         "Non-valid URI path",
// 			RequestURI:   "qwerty",
// 			ExpectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			Name:         "Non-valid filter",
// 			RequestURI:   fmt.Sprintf("%d", bookID),
// 			RequestQuery: `page=none,page_size=many,sort=idk`,
// 			ExpectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			Name:           "Error while retriving",
// 			RequestURI:     fmt.Sprintf("%d", bookID),
// 			RequestQuery:   `page=1&page_size=10&sort=created_at`,
// 			MockResult:     nil,
// 			MockResultMeta: nil,
// 			MockResultErr:  errors.New("critical error"),
// 			ExpectedID:     bookID,
// 			ExpectedCode:   http.StatusInternalServerError,
// 		},
// 	}

// 	var userID int64 = 5
// 	var bookID int64 = 42
// 	tests := []struct {
// 		Name           string
// 		MockResult     any
// 		ExpectedReview *entity.Review
// 	}{
// 		{
// 			Name:       "Review updated successfullt",
// 			MockResult: nil,
// 			ExpectedReview: &entity.Review{
// 				Content: util.StringToPointer("Very good"),
// 				Rating:  util.IntToPointer(100),
// 				UserID:  userID,
// 				BookID:  bookID,
// 			},
// 		},
// 		{
// 			Name:       "Some error ocurred while saving",
// 			MockResult: errors.New("criticabookIDl error"),
// 			ExpectedReview: &entity.Review{
// 				Content: util.StringToPointer("Very good"),
// 				Rating:  util.IntToPointer(100),
// 				UserID:  userID,
// 				BookID:  bookID,
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			repo := mocks.NewRepository(t)
// 			service := New(repo, nil)
// 			ctx := context.Background()

// 			repo.On("UpdateReview", ctx, test.ExpectedReview).Return(test.MockResult)

// 			assert.Equal(t, service.UpdateReview(ctx, test.ExpectedReview), test.MockResult)
// 		})
// 	}
// 	repo := mocks.NewRepository(t)
// 	service := New(repo, nil)
// 	ctx := context.Background()

// 	var id int64 = 5
// 	reviews := make([]*entity.Review, 5)

// 	filter := util.NewFilter(1, 50, "created_at")
// 	invalidFilter := util.NewFilter(1, 50, "name")

// 	meta := filter.CalculateMetadata(len(reviews))

// 	repo.On("GetReviewsForBookID", ctx, id, filter).Return(reviews, &meta, nil).Once()
// 	repo.On("GetReviewsForBookID", ctx, id, filter).Return(nil, nil, errors.New("some error")).Once()

// 	result, _, err := service.GetReviewsForBookID(ctx, id, filter)
// 	assert.NotNil(t, result)
// 	assert.Nil(t, err)

// 	result, _, err = service.GetReviewsForBookID(ctx, id, invalidFilter)
// 	assert.Nil(t, result)
// 	assert.NotNil(t, err)

// 	result, _, err = service.GetReviewsForBookID(ctx, id, filter)
// 	assert.Nil(t, result)
// 	assert.NotNil(t, err)
// }

func TestUpdateReview(t *testing.T) {
	var userID int64 = 5
	var bookID int64 = 42
	tests := []struct {
		Name           string
		MockResult     any
		ExpectedReview *entity.Review
	}{
		{
			Name:       "Review updated successfullt",
			MockResult: nil,
			ExpectedReview: &entity.Review{
				Content: util.StringToPointer("Very good"),
				Rating:  util.IntToPointer(100),
				UserID:  userID,
				BookID:  bookID,
			},
		},
		{
			Name:       "Some error ocurred while saving",
			MockResult: errors.New("criticabookIDl error"),
			ExpectedReview: &entity.Review{
				Content: util.StringToPointer("Very good"),
				Rating:  util.IntToPointer(100),
				UserID:  userID,
				BookID:  bookID,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("UpdateReview", ctx, test.ExpectedReview).Return(test.MockResult)

			assert.Equal(t, service.UpdateReview(ctx, test.ExpectedReview), test.MockResult)
		})
	}
}

func TestDeleteReview(t *testing.T) {
	var userID int64 = 5
	var reviewID int64 = 42
	tests := []struct {
		Name       string
		MockResult any
	}{
		{
			Name:       "Review deleted successfullt",
			MockResult: nil,
		},
		{
			Name:       "Some error ocurred while deleting",
			MockResult: errors.New("critical error"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("DeleteReview", ctx, reviewID, userID).Return(test.MockResult)

			assert.Equal(t, service.DeleteReview(ctx, reviewID, userID), test.MockResult)
		})
	}
}
