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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSuspendUser(t *testing.T) {
	var userID int64 = 123
	expiresIn := time.Minute * 100
	tests := []struct {
		Name               string
		RequestJSON        string
		MockResult         any
		ExpectedSuspension *entity.Suspension
		ExpectedCode       int
	}{
		{
			Name: "Suspend user successfully",
			RequestJSON: `{
				"reason": "Bad behaviour",
				"expires_in": 100,
				"user_id": 2
			}`,
			MockResult: nil,
			ExpectedSuspension: &entity.Suspension{
				UserID:      2,
				Reason:      util.StringToPointer("Bad behaviour"),
				ModeratorID: userID,
				ExpiresIn:   &expiresIn,
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Name: "Non-valid json",
			RequestJSON: `{
				"expires_in": 100,
				"user_id": 2
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "Error while saving record",
			RequestJSON: `{
				"reason": "Bad behaviour",
				"expires_in": 100,
				"user_id": 2
			}`,
			MockResult: errors.New("critical error"),
			ExpectedSuspension: &entity.Suspension{
				UserID:      2,
				Reason:      util.StringToPointer("Bad behaviour"),
				ModeratorID: userID,
				ExpiresIn:   &expiresIn,
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

			req, _ := http.NewRequest("POST", "/mod/suspensions/new", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req
			ctx.Set("userID", userID)

			service.On("NewSuspension", ctx, test.ExpectedSuspension).Return(test.MockResult)
			handler.suspendUser(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}

func TestCheckSuspension(t *testing.T) {
	var userID int64 = 123
	expiresIn := time.Minute * 100
	tests := []struct {
		Name         string
		RequestURI   string
		MockResult   any
		MockError    error
		ExpectedID   int64
		ExpectedCode int
	}{
		{
			Name:       "Get suspensions succesfully",
			RequestURI: fmt.Sprintf("%d", userID),
			MockResult: []*entity.Suspension{{
				UserID:      2,
				Reason:      util.StringToPointer("Bad behaviour"),
				ModeratorID: userID,
				ExpiresIn:   &expiresIn,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			}},
			ExpectedID:   userID,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid id",
			RequestURI:   "bleh",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "Error while retrieving records",
			RequestURI:   fmt.Sprintf("%d", userID),
			MockResult:   nil,
			MockError:    errors.New("critical error"),
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

			req, _ := http.NewRequest("GET", "/users/suspensions/"+test.Name, strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req

			service.On("CheckSuspension", ctx, test.ExpectedID).Return(test.MockResult, test.MockError)
			handler.checkSuspension(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}

}

func TestUpdateSuspension(t *testing.T) {
	var userID int64 = 123
	expiresIn := time.Minute * 100
	tests := []struct {
		Name               string
		RequestURI         string
		RequestJSON        string
		MockResult         any
		ExpectedSuspension *entity.Suspension
		ExpectedCode       int
	}{
		{
			Name:       "Update suspension successfully",
			RequestURI: "2",
			RequestJSON: `{
				"reason": "Bad behaviour",
				"expires_in": 100
			}`,
			MockResult: nil,
			ExpectedSuspension: &entity.Suspension{
				ID:          2,
				Reason:      util.StringToPointer("Bad behaviour"),
				ModeratorID: userID,
				ExpiresIn:   &expiresIn,
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Non-valid id",
			RequestURI:   "bleh",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Non-valid json",
			RequestURI: "2",
			RequestJSON: `{
				bleh
			}`,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "Error while saving record",
			RequestURI: "2",
			RequestJSON: `{
				"reason": "Bad behaviour",
				"expires_in": 100
			}`,
			MockResult: errors.New("critical error"),
			ExpectedSuspension: &entity.Suspension{
				ID:          2,
				Reason:      util.StringToPointer("Bad behaviour"),
				ModeratorID: userID,
				ExpiresIn:   &expiresIn,
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

			req, _ := http.NewRequest("POST", "/mod/suspensions/new", strings.NewReader(test.RequestJSON))
			req.Header.Set("Content-Type", "application/json")

			param := gin.Param{Key: "id", Value: test.RequestURI}
			ctx.Params = append(ctx.Params, param)
			ctx.Request = req
			ctx.Set("userID", userID)

			service.On("UpdateSuspension", ctx, test.ExpectedSuspension).Return(test.MockResult)
			handler.updateSuspension(ctx)
			assert.Equal(t, test.ExpectedCode, w.Code)
		})
	}
}
