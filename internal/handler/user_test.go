package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/service/mocks"
	"one-lab-final/pkg/util"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		Name         string
		RequestJSON  string
		MockResult   any
		ExpectedUser entity.User
		ExpectedCode int
	}{
		{
			Name: "Create user successfully",
			RequestJSON: `
			{
				"username": "Flove",
				"email": "example@gmail.com",
				"password": "password",
				"first_name": "Firstname",
				"last_name": "Surname"
			}`,
			MockResult: nil,
			ExpectedUser: entity.User{
				Username:  util.StringToPointer("Flove"),
				Email:     util.StringToPointer("example@gmail.com"),
				FirstName: util.StringToPointer("Firstname"),
				LastName:  util.StringToPointer("Surname"),
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Name: "Empty fields",
			RequestJSON: `
			{
				"email": "example@gmail.com",
				"first_name": "Firstname",
				"last_name": "Surname"
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Non-valid JSON",
			RequestJSON: `
			{
				Boo
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Error while saving entity",
			RequestJSON: `
			{
				"username": "Flove",
				"email": "example@gmail.com",
				"password": "password",
				"first_name": "Firstname",
				"last_name": "Surname"
			}`,
			ExpectedUser: entity.User{
				Username:  util.StringToPointer("Flove"),
				Email:     util.StringToPointer("example@gmail.com"),
				FirstName: util.StringToPointer("Firstname"),
				LastName:  util.StringToPointer("Surname"),
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			MockResult:   errors.New("critical error"),
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

			req, _ := http.NewRequest("POST", "/users/new", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req

			service.On("CreateUser", ctx, &test.ExpectedUser).Return(test.MockResult)
			handler.createUser(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		Name                string
		RequestJSON         string
		MockResult          any
		MockError           error
		ExpectedCredentials string
		ExpectedPassword    string
		ExpectedCode        int
	}{
		{
			Name: "Logged in successfully",
			RequestJSON: `
			{
				"credentials": "flove",
				"password": "password"
			}`,
			MockResult:          &entity.Token{},
			ExpectedCredentials: "flove",
			ExpectedPassword:    "password",
			ExpectedCode:        http.StatusCreated,
		},
		{
			Name: "Non-valid JSON",
			RequestJSON: `
			{
				Boo
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Mismatching password",
			RequestJSON: `
			{
				"credentials": "flove",
				"password": "password"
			}`,
			MockError:           util.ErrMismatchedPassword,
			ExpectedCredentials: "flove",
			ExpectedPassword:    "password",
			ExpectedCode:        http.StatusUnauthorized,
		},
		{
			Name: "Error while saving entity",
			RequestJSON: `
			{
				"credentials": "flove",
				"password": "password"
			}`,
			MockError:           errors.New("critical error"),
			ExpectedCredentials: "flove",
			ExpectedPassword:    "password",
			ExpectedCode:        http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("POST", "/users/login", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req

			service.On("Login", ctx, test.ExpectedCredentials, test.ExpectedPassword).Return(test.MockResult, test.MockError)
			handler.login(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	var userID int64 = 123
	tests := []struct {
		Name         string
		RequestJSON  string
		MockResult   any
		ExpectedUser entity.User
		ExpectedCode int
	}{
		{
			Name: "Update user fully",
			RequestJSON: `
			{
				"password": "password",
				"first_name": "Firstname",
				"last_name": "Surname"
			}`,
			ExpectedUser: entity.User{
				ID:        userID,
				FirstName: util.StringToPointer("Firstname"),
				LastName:  util.StringToPointer("Surname"),
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			MockResult:   nil,
			ExpectedCode: http.StatusOK,
		},
		{
			Name: "Update user partialy",
			RequestJSON: `
			{
				"first_name": "Firstname"
			}`,
			MockResult: nil,
			ExpectedUser: entity.User{
				ID:        userID,
				FirstName: util.StringToPointer("Firstname"),
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid URI path",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Non-valid json body",
			RequestJSON: `
			{
				OOPS
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Error while saving entity",
			RequestJSON: `{
				"password": "password",
				"first_name": "Firstname",
				"last_name": "Surname"
			}`,
			ExpectedUser: entity.User{
				ID:        userID,
				FirstName: util.StringToPointer("Firstname"),
				LastName:  util.StringToPointer("Surname"),
				Password: entity.Password{
					Plaintext: util.StringToPointer("password"),
				},
			},
			MockResult:   errors.New("critical error"),
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

			req, _ := http.NewRequest("PATCH", "/users/update", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Set("userID", userID)
			ctx.Request = req

			service.On("UpdateUser", ctx, &test.ExpectedUser).Return(test.MockResult)
			handler.updateUser(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	var userID int64 = 123
	tests := []struct {
		Name         string
		MockResult   any
		ExpectedID   int64
		ExpectedCode int
	}{
		{
			Name:         "Delete user successfully",
			MockResult:   nil,
			ExpectedID:   userID,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Error while deleting entity",
			MockResult:   errors.New("critical error"),
			ExpectedID:   userID,
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

			req, _ := http.NewRequest("DELETE", "/user/delete", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			ctx.Set("userID", userID)
			ctx.Request = req

			service.On("DeleteUser", ctx, test.ExpectedID).Return(test.MockResult)
			handler.deleteUser(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
