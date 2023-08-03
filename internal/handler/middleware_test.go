package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository"
	"one-lab-final/internal/service/mocks"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		Name           string
		MockResult     any
		MockError      error
		Token          string
		ExpectedToken  string
		ExpectedCode   int
		ExpectedStatus bool
	}{
		{
			Name:           "Authenticated succesfully",
			MockResult:     &entity.User{},
			Token:          "Bearer token",
			ExpectedToken:  "token",
			ExpectedCode:   http.StatusOK,
			ExpectedStatus: true,
		},
		{
			Name:           "Non-valid token",
			MockError:      repository.ErrRecordNotFound,
			Token:          "Bearer token",
			ExpectedToken:  "token",
			ExpectedCode:   http.StatusUnauthorized,
			ExpectedStatus: false,
		},
		{
			Name:           "Missing header",
			Token:          "",
			ExpectedCode:   http.StatusUnauthorized,
			ExpectedStatus: false,
		},
		{
			Name:           "Malformed header",
			Token:          "Bearer",
			ExpectedCode:   http.StatusUnauthorized,
			ExpectedStatus: false,
		},
		{
			Name:           "Error while retrieving record",
			MockError:      errors.New("critical error"),
			Token:          "Bearer token",
			ExpectedToken:  "token",
			ExpectedCode:   http.StatusInternalServerError,
			ExpectedStatus: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			service := &mocks.Service{}
			handler := New(service, nil)

			req, _ := http.NewRequest("POST", "/books/new", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", test.Token)

			ctx.Request = req

			service.On("GetUserByToken", ctx, test.ExpectedToken).Return(test.MockResult, test.MockError)
			status := handler.authenticate(ctx)

			assert.Equal(t, test.ExpectedStatus, status)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestRequireRole(t *testing.T) {
	tests := []struct {
		Name          string
		MockResult    any
		MockError     error
		Token         string
		ExpectedToken string
		ExpectedRole  entity.Role
		ExpectedCode  int
	}{
		{
			Name: "Authenticated succesfully",
			MockResult: &entity.User{
				Role: entity.ADMIN,
			},
			Token:         "Bearer token",
			ExpectedToken: "token",
			ExpectedRole:  entity.ADMIN,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:         "Non-valid token",
			Token:        "Bearer",
			ExpectedCode: http.StatusUnauthorized,
		},
		{
			Name: "Roles do not match",
			MockResult: &entity.User{
				Role: entity.USER,
			},
			Token:         "Bearer token",
			ExpectedToken: "token",
			ExpectedRole:  entity.ADMIN,
			ExpectedCode:  http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, e := gin.CreateTestContext(w)
			service := &mocks.Service{}
			handler := New(service, nil)

			e.Use(handler.requireRole(test.ExpectedRole)).GET("/healthcheck", handler.healthcheck)
			req, _ := http.NewRequest("GET", "/healthcheck", strings.NewReader(""))
			req.Header.Set("Authorization", test.Token)
			ctx.Request = req

			service.On("GetUserByToken", mock.AnythingOfType("*gin.Context"), test.ExpectedToken).Return(test.MockResult, test.MockError)
			e.ServeHTTP(w, req)

			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
func TestRequireAuthenticatedUser(t *testing.T) {
	tests := []struct {
		Name          string
		MockResult    any
		MockError     error
		Token         string
		ExpectedToken string
		ExpectedRole  entity.Role
		ExpectedCode  int
	}{
		{
			Name: "Authenticated succesfully",
			MockResult: &entity.User{
				Role: entity.ADMIN,
			},
			Token:         "Bearer token",
			ExpectedToken: "token",
			ExpectedRole:  entity.ADMIN,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:         "Non-valid token",
			Token:        "Bearer",
			ExpectedCode: http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, e := gin.CreateTestContext(w)
			service := &mocks.Service{}
			handler := New(service, nil)

			e.Use(handler.requireAuthenticatedUser()).GET("/healthcheck", handler.healthcheck)
			req, _ := http.NewRequest("GET", "/healthcheck", strings.NewReader(""))
			req.Header.Set("Authorization", test.Token)
			ctx.Request = req

			service.On("GetUserByToken", mock.AnythingOfType("*gin.Context"), test.ExpectedToken).Return(test.MockResult, test.MockError)
			e.ServeHTTP(w, req)

			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
