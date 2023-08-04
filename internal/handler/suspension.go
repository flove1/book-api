package handler

import (
	"fmt"
	"net/http"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/handler/api"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Suspend user for some time. Requires MODERATOR role or higher
// @Tags         Moderation
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param data body api.CreateSuspensionRequest true "Request body"
//
// @Success      200 {object} api.DefaultResponse "User was successfully suspended"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /mod/suspensions/new [post]
func (h *Handler) suspendUser(ctx *gin.Context) {
	var req api.CreateSuspensionRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ExpiresIn := time.Minute * time.Duration(req.ExpiresIn)

	suspension := &entity.Suspension{
		Reason:      &req.Reason,
		UserID:      req.UserID,
		ModeratorID: ctx.MustGet("userID").(int64),
		ExpiresIn:   &ExpiresIn,
	}

	err = h.Services.NewSuspension(ctx, suspension)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Header("Locations", fmt.Sprintf("/mod/suspensions/%d", suspension.ID))

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "user was successfully suspended",
	})
}

// @Summary      Check if user has suspensions
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
//
// @Success      200 {object} api.CheckSuspensionResponse "ok"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/suspensions/{id} [get]
func (h *Handler) checkSuspension(ctx *gin.Context) {
	var req api.ID

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	suspensions, err := h.Services.CheckSuspension(ctx, req.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.CheckSuspensionResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    suspensions,
	})
}

// @Summary      Modify suspension of user. Requires MODERATOR role or higher
// @Tags         Moderation
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Suspension ID"
// @Param data body api.UpdateSuspensionRequest true "Request body"
//
// @Success      200 {object} api.DefaultResponse "suspension was successfully updated"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /mod/suspensions/update/{id} [patch]
func (h *Handler) updateSuspension(ctx *gin.Context) {
	var req api.UpdateSuspensionRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
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

	userID := ctx.MustGet("userID").(int64)
	ExpiresIn := time.Minute * time.Duration(req.ExpiresIn)

	err = h.Services.UpdateSuspension(ctx, &entity.Suspension{
		ID:          req.ID.Value,
		Reason:      &req.Reason,
		ModeratorID: userID,
		ExpiresIn:   &ExpiresIn,
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
		Message: "suspension was successfully updated",
	})
}
