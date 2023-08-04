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

func TestCreateUser(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		User       *entity.User
		ExpectErr  bool
	}{
		{
			Name:       "User created successfully",
			MockResult: nil,
			User: &entity.User{
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
		},
		{
			Name:       "Some error ocurred while saving",
			MockResult: errors.New("critical error"),
			User: &entity.User{
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			ExpectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("CreateUser", ctx, test.User).Return(test.MockResult)

			err := service.CreateUser(ctx, test.User)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetUserByCredentials(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		MockError  error
		Username   string
		ExpectErr  bool
	}{
		{
			Name:       "User exists",
			MockResult: nil,
			Username:   "username",
		},
		{
			Name:      "Some error ocurred while retrieving",
			MockError: errors.New("critical error"),
			Username:  "username",
			ExpectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("GetUserByCredentials", ctx, test.Username).Return(test.MockResult, test.MockError)

			_, err := service.GetUserByCredentials(ctx, test.Username)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetUserByToken(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		MockError  error
		Token      string
		ExpectErr  bool
	}{
		{
			Name:       "User exists",
			MockResult: nil,
			Token:      "gsrdnlkcnjsq213nj12",
		},
		{
			Name:      "Some error ocurred while retrieving",
			MockError: errors.New("critical error"),
			Token:     "gsrdnlkcnjsq213nj12",
			ExpectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("GetUserByToken", ctx, test.Token).Return(test.MockResult, test.MockError)

			_, err := service.GetUserByToken(ctx, test.Token)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		User       *entity.User
		ExpectErr  bool
	}{
		{
			Name:       "User updated successfully",
			MockResult: nil,
			User: &entity.User{
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
		},
		{
			Name:       "Some error ocurred while updating",
			MockResult: errors.New("critical error"),
			User: &entity.User{
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			ExpectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("UpdateUser", ctx, test.User).Return(test.MockResult)

			err := service.UpdateUser(ctx, test.User)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		UserID     int64
		ExpectErr  bool
	}{
		{
			Name:       "User deleted successfully",
			MockResult: nil,
			UserID:     10,
		},
		{
			Name:       "Some error ocurred while deleting",
			MockResult: errors.New("critical error"),
			UserID:     10,
			ExpectErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("DeleteUser", ctx, test.UserID).Return(test.MockResult)

			err := service.DeleteUser(ctx, test.UserID)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

// func TestGetReviewsForBookID(t *testing.T) {
// 	var bookID int64 = 123
// 	tests := []struct {
// 		Name           string
// 		MockResultErr  error
// 		MockResultMeta *util.Metadata
// 		MockResult     any
// 		SkipMock       bool
// 		BookID         int64
// 		Filter         util.Filter
// 		ExpectedErr    bool
// 	}{
// 		{
// 			Name: "Get reviews succesfuly",
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
// 			BookID:         bookID,
// 			Filter:         util.NewFilter(1, 10, "created_at"),
// 			ExpectedErr:    false,
// 		},
// 		{
// 			Name:        "Non-valid sort",
// 			SkipMock:    true,
// 			Filter:      util.NewFilter(1, 10, "everything"),
// 			ExpectedErr: true,
// 		},
// 		{
// 			Name:          "Error while retriving",
// 			BookID:        bookID,
// 			MockResultErr: errors.New("critical error"),
// 			Filter:        util.NewFilter(1, 10, "created_at"),
// 			ExpectedErr:   true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			repo := mocks.NewRepository(t)
// 			service := New(repo, nil)
// 			ctx := context.Background()

// 			if !test.SkipMock {
// 				repo.On("GetReviewsByBookID", ctx, test.BookID, test.Filter).Return(test.MockResult, test.MockResultMeta, test.MockResultErr)
// 			}

// 			_, _, err := service.GetReviewsByBookID(ctx, test.BookID, test.Filter)

// 			if test.ExpectedErr {
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

// func TestUpdateReview(t *testing.T) {
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
// }

// func TestDeleteReview(t *testing.T) {
// 	var userID int64 = 5
// 	var reviewID int64 = 42
// 	tests := []struct {
// 		Name       string
// 		MockResult any
// 	}{
// 		{
// 			Name:       "Review deleted successfullt",
// 			MockResult: nil,
// 		},
// 		{
// 			Name:       "Some error ocurred while deleting",
// 			MockResult: errors.New("critical error"),
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			repo := mocks.NewRepository(t)
// 			service := New(repo, nil)
// 			ctx := context.Background()

// 			repo.On("DeleteReview", ctx, reviewID, userID).Return(test.MockResult)

// 			assert.Equal(t, service.DeleteReview(ctx, reviewID, userID), test.MockResult)
// 		})
// 	}
