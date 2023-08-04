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

func TestGetUserByUsername(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		MockError  error
		Username   string
		ExpectErr  bool
	}{
		{
			Name:       "User exists",
			MockResult: &entity.User{},
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

			repo.On("GetUserByUsername", ctx, test.Username).Return(test.MockResult, test.MockError)

			_, err := service.GetUserByUsername(ctx, test.Username)

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
		Name        string
		MockResult  any
		MockError   error
		Credentials string
		ExpectErr   bool
	}{
		{
			Name:        "Get by email",
			MockResult:  &entity.User{},
			Credentials: "example@gmail.com",
		},
		{
			Name:        "Get by username",
			MockResult:  &entity.User{},
			Credentials: "username",
		},
		{
			Name:        "Some error ocurred while retrieving",
			MockError:   errors.New("critical error"),
			Credentials: "username",
			ExpectErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("GetUserByCredentials", ctx, test.Credentials).Return(test.MockResult, test.MockError)

			_, err := service.GetUserByCredentials(ctx, test.Credentials)

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

func TestGrantRoleToUser(t *testing.T) {
	var userID int64 = 15
	tests := []struct {
		Name       string
		MockResult any
		UserID     int64
		UserRole   entity.Role
		ExpectErr  bool
	}{
		{
			Name:       "Role updated successfully",
			MockResult: nil,
			UserID:     userID,
			UserRole:   entity.ADMIN,
		},
		{
			Name:       "Some error ocurred while updating",
			MockResult: errors.New("critical error"),
			UserID:     userID,
			UserRole:   entity.ADMIN,
			ExpectErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("GrantRoleToUser", ctx, test.UserID, test.UserRole).Return(test.MockResult)

			err := service.GrantRoleToUser(ctx, test.UserID, test.UserRole)

			if test.ExpectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
