package api

import (
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

type LoginResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Body    *entity.Token `json:"body"`
}

type GetUserByUsernameResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Body    *entity.User `json:"body"`
}

type GetBookByIDResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Body    *entity.Book `json:"body"`
}

type GetBooksResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.Book `json:"body"`
	Meta    util.Metadata  `json:"meta"`
}

type GetReviewsByBookIDResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Body    []*entity.Review `json:"body"`
	Meta    util.Metadata    `json:"meta"`
}

type CheckSuspensionResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Body    []*entity.Suspension `json:"body"`
}
