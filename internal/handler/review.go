package handler

import (
	"errors"
	"fmt"
	"net/http"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/handler/api"
	"one-lab-final/internal/repository"
	"one-lab-final/pkg/util"

	"github.com/gin-gonic/gin"
)

// @Summary      Create new review
// @Tags         Reviews
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param data body api.CreateReviewRequest true "Request body"
//
// @Success      201 {object} api.DefaultResponse "Review succesfully created"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /reviews/new [post]
func (h *Handler) createReview(ctx *gin.Context) {
	var req api.CreateReviewRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.UserID = ctx.MustGet("userID").(int64)

	err = h.Services.CreateReview(ctx, &entity.Review{
		Content: &req.Content,
		Rating:  &req.Rating,
		UserID:  req.UserID,
		BookID:  req.BookID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "review succesfully created",
	})
}

// @Summary      Get reviews by book ID using pagination
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param filter  query api.Filter true "Pagination filter"
//
// @Success      200 {object} api.GetReviewsByBookIDResponse "Reviews info"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/{id}/reviews [get]
func (h *Handler) getReviewsByBookID(ctx *gin.Context) {
	var req api.GetReviewsByBookIDRequest
	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	reviews, meta, err := h.Services.GetReviewsByBookID(ctx, id.Value, util.NewFilter(req.Page, req.PageSize, req.Sort))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, &api.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "book does not exists",
			})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, &api.GetReviewsByBookIDResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    reviews,
		Meta:    *meta,
	})
}

// @Summary      Update review by ID
// @Tags         Reviews
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Review ID"
// @Param data body api.UpdateReviewRequest true "Request body"
//
// @Success      200 {object} api.DefaultResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /reviews/update/{id} [patch]
func (h *Handler) updateReview(ctx *gin.Context) {
	var req api.UpdateReviewRequest
	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.UserID = ctx.MustGet("userID").(int64)

	err = h.Services.UpdateReview(ctx, &entity.Review{
		ID:      id.Value,
		Content: req.Content,
		Rating:  req.Rating,
		UserID:  req.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "review succesfully updated",
	})

}

// @Summary      Delete review by ID
// @Tags         Reviews
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Review ID"
//
// @Success      200 {object} api.DefaultResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /reviews/delete/{id} [delete]
func (h *Handler) deleteReview(ctx *gin.Context) {
	var req api.DeleteReviewRequest
	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.UserID = ctx.MustGet("userID").(int64)

	err = h.Services.DeleteReview(ctx, id.Value, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// if req.Role >= entity.MODERATOR {
	// 	log.Println()
	// }

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "review succesfully deleted",
	})
}
