package api

import "one-lab-final/internal/entity"

type CreateReviewRequest struct {
	UserID  int64  `json:"-"`
	Content string `json:"content" binding:"required" example:"Awesome book"`
	Rating  int64  `json:"rating,default=0" binding:"required" example:"100" default:"0"`
	BookID  int64  `json:"book_id" example:"1"`
}

type UpdateReviewRequest struct {
	UserID  int64   `json:"-"`
	Content *string `json:"content" binding:"omitempty" example:"Awesome book"`
	Rating  *int64  `json:"rating" binding:"omitempty" example:"100" default:"0"`
}

type GetReviewsByBookIDRequest struct {
	Filter
}

type DeleteReviewRequest struct {
	UserID int64
	Role   entity.Role
}
