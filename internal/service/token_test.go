package service

import (
	"context"
	"errors"
	"one-lab-final/internal/config"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository"
	"one-lab-final/internal/repository/mocks"
	"one-lab-final/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	password := "password"
	hash, _ := util.HashPassword(password)
	tests := []struct {
		Name            string
		MockUser        *entity.User
		MockUserErr     error
		MockTokenResult error
		Credentials     string
		Password        string
		ExpectErr       bool
	}{
		{
			Name:        "Token created successfully",
			Credentials: "username",
			Password:    password,
			MockUser: &entity.User{
				Password: entity.Password{
					Plaintext: &password,
					Hash:      &hash,
				},
			},
			MockTokenResult: nil,
		},
		{
			Name:        "User doesn't exist",
			Credentials: "username",
			Password:    password,
			MockUser:    nil,
			MockUserErr: repository.ErrRecordNotFound,
			ExpectErr:   true,
		},
		{
			Name:        "Some error ocurred while saving token",
			Credentials: "username",
			Password:    password,
			MockUser: &entity.User{
				Password: entity.Password{
					Plaintext: &password,
					Hash:      &hash,
				},
			},
			MockTokenResult: errors.New("critical error"),
			ExpectErr:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, &config.Config{
				AUTH: config.AuthConfig{
					SigningKey: "BLEH",
				},
			})
			ctx := context.Background()

			repo.On("GetUserByCredentials", ctx, test.Credentials).Return(test.MockUser, test.MockUserErr)
			if test.MockUserErr == nil {
				repo.On("CreateToken", ctx, mock.AnythingOfType("*entity.Token")).Return(test.MockTokenResult)
			}

			_, err := service.Login(ctx, test.Credentials, test.Password)
			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteExpiredTokens(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
	}{
		{
			Name:       "Deleted successfully",
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

			repo.On("DeleteExpiredTokens", ctx).Return(test.MockResult)

			assert.Equal(t, service.DeleteExpiredTokens(ctx), test.MockResult)
		})
	}
}
