package service

import (
	"context"
	"errors"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuspension(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		Suspension *entity.Suspension
	}{
		{
			Name:       "User successfully suspended",
			MockResult: nil,
			Suspension: &entity.Suspension{},
		},
		{
			Name:       "Some error ocurred while saving",
			MockResult: errors.New("critical error"),
			Suspension: &entity.Suspension{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("NewSuspension", ctx, test.Suspension).Return(test.MockResult)

			assert.Equal(t, service.NewSuspension(ctx, test.Suspension), test.MockResult)
		})
	}
}

func TestCheckSuspension(t *testing.T) {
	var userID int64 = 123
	tests := []struct {
		Name        string
		MockResult  any
		MockError   error
		ExpectError bool
	}{
		{
			Name:        "Suspensions retrieved successfully",
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

			repo.On("CheckSuspension", ctx, userID).Return(test.MockResult, test.MockError)

			_, err := service.CheckSuspension(ctx, userID)

			if test.ExpectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateSuspension(t *testing.T) {
	tests := []struct {
		Name       string
		MockResult any
		Suspension *entity.Suspension
	}{
		{
			Name:       "Suspension updated successfully",
			MockResult: nil,
			Suspension: &entity.Suspension{},
		},
		{
			Name:       "Some error ocurred while updating",
			MockResult: errors.New("critical error"),
			Suspension: &entity.Suspension{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			service := New(repo, nil)
			ctx := context.Background()

			repo.On("UpdateSuspension", ctx, test.Suspension).Return(test.MockResult)

			assert.Equal(t, service.UpdateSuspension(ctx, test.Suspension), test.MockResult)
		})
	}
}
