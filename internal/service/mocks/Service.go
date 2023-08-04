// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "one-lab-final/internal/entity"

	mock "github.com/stretchr/testify/mock"

	util "one-lab-final/pkg/util"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckSuspension provides a mock function with given fields: ctx, userID
func (_m *Service) CheckSuspension(ctx context.Context, userID int64) ([]*entity.Suspension, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*entity.Suspension
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*entity.Suspension, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entity.Suspension); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Suspension)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBook provides a mock function with given fields: ctx, book
func (_m *Service) CreateBook(ctx context.Context, book *entity.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateReview provides a mock function with given fields: ctx, review
func (_m *Service) CreateReview(ctx context.Context, review *entity.Review) error {
	ret := _m.Called(ctx, review)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Review) error); ok {
		r0 = rf(ctx, review)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Service) CreateUser(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: ctx, articleID
func (_m *Service) DeleteBook(ctx context.Context, articleID int64) error {
	ret := _m.Called(ctx, articleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, articleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExpiredTokens provides a mock function with given fields: ctx
func (_m *Service) DeleteExpiredTokens(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReview provides a mock function with given fields: ctx, reviewID, userID
func (_m *Service) DeleteReview(ctx context.Context, reviewID int64, userID int64) error {
	ret := _m.Called(ctx, reviewID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, reviewID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Service) DeleteUser(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBookByID provides a mock function with given fields: ctx, bookID
func (_m *Service) GetBookByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	ret := _m.Called(ctx, bookID)

	var r0 *entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Book, error)); ok {
		return rf(ctx, bookID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Book); ok {
		r0 = rf(ctx, bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: ctx, title, author, tags, filter
func (_m *Service) GetBooks(ctx context.Context, title *string, author *string, tags *[]string, filter util.Filter) ([]*entity.Book, *util.Metadata, error) {
	ret := _m.Called(ctx, title, author, tags, filter)

	var r0 []*entity.Book
	var r1 *util.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, *string, *[]string, util.Filter) ([]*entity.Book, *util.Metadata, error)); ok {
		return rf(ctx, title, author, tags, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string, *string, *[]string, util.Filter) []*entity.Book); ok {
		r0 = rf(ctx, title, author, tags, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string, *string, *[]string, util.Filter) *util.Metadata); ok {
		r1 = rf(ctx, title, author, tags, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*util.Metadata)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *string, *string, *[]string, util.Filter) error); ok {
		r2 = rf(ctx, title, author, tags, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetReviewsByBookID provides a mock function with given fields: ctx, bookID, filter
func (_m *Service) GetReviewsByBookID(ctx context.Context, bookID int64, filter util.Filter) ([]*entity.Review, *util.Metadata, error) {
	ret := _m.Called(ctx, bookID, filter)

	var r0 []*entity.Review
	var r1 *util.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, util.Filter) ([]*entity.Review, *util.Metadata, error)); ok {
		return rf(ctx, bookID, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, util.Filter) []*entity.Review); ok {
		r0 = rf(ctx, bookID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Review)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, util.Filter) *util.Metadata); ok {
		r1 = rf(ctx, bookID, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*util.Metadata)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, util.Filter) error); ok {
		r2 = rf(ctx, bookID, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserByCredentials provides a mock function with given fields: ctx, credentials
func (_m *Service) GetUserByCredentials(ctx context.Context, credentials string) (*entity.User, error) {
	ret := _m.Called(ctx, credentials)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, credentials)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, credentials)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentials)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByToken provides a mock function with given fields: ctx, token
func (_m *Service) GetUserByToken(ctx context.Context, token string) (*entity.User, error) {
	ret := _m.Called(ctx, token)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *Service) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GrantRoleToUser provides a mock function with given fields: ctx, userID, role
func (_m *Service) GrantRoleToUser(ctx context.Context, userID int64, role entity.Role) error {
	ret := _m.Called(ctx, userID, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, entity.Role) error); ok {
		r0 = rf(ctx, userID, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, credentials, password
func (_m *Service) Login(ctx context.Context, credentials string, password string) (*entity.Token, error) {
	ret := _m.Called(ctx, credentials, password)

	var r0 *entity.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*entity.Token, error)); ok {
		return rf(ctx, credentials, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entity.Token); ok {
		r0 = rf(ctx, credentials, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, credentials, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSuspension provides a mock function with given fields: ctx, suspension
func (_m *Service) NewSuspension(ctx context.Context, suspension *entity.Suspension) error {
	ret := _m.Called(ctx, suspension)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Suspension) error); ok {
		r0 = rf(ctx, suspension)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshBooksRating provides a mock function with given fields: ctx
func (_m *Service) RefreshBooksRating(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBook provides a mock function with given fields: ctx, book
func (_m *Service) UpdateBook(ctx context.Context, book *entity.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReview provides a mock function with given fields: ctx, review
func (_m *Service) UpdateReview(ctx context.Context, review *entity.Review) error {
	ret := _m.Called(ctx, review)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Review) error); ok {
		r0 = rf(ctx, review)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSuspension provides a mock function with given fields: ctx, suspension
func (_m *Service) UpdateSuspension(ctx context.Context, suspension *entity.Suspension) error {
	ret := _m.Called(ctx, suspension)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Suspension) error); ok {
		r0 = rf(ctx, suspension)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *Service) UpdateUser(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
