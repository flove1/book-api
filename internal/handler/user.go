package handler

import (
	"errors"
	"net/http"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/handler/api"
	"one-lab-final/internal/repository"
	"one-lab-final/pkg/util"

	"github.com/gin-gonic/gin"
)

// @Summary      Register new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param data body api.CreateUserRequest true "Request body"
//
// @Success      201 {object} api.DefaultResponse "User succesfully created"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /user/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.Services.CreateUser(ctx, &entity.User{
		Username:  &req.Username,
		Email:     &req.Email,
		FirstName: &req.FirstName,
		LastName:  &req.LastName,
		Password: entity.Password{
			Plaintext: &req.Password,
		},
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
		Message: "user succesfully created",
	})
}

// @Summary      Get user by his username
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        username   path      string  true  "Username of user"
//
// @Success      200 {object} api.GetUserByUsernameResponse "Ok"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /user/{username} [get]
func (h *Handler) getUserByUsername(ctx *gin.Context) {
	var req api.Username

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := h.Services.GetUserByCredentials(ctx, req.Value)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, &api.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "user does not exists",
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

	ctx.JSON(http.StatusOK, &api.GetUserByUsernameResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    user,
	})
}

// @Summary      Authenticated user and return access token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param data body api.LoginRequest true "Request body"
//
// @Success      201 {object} api.LoginResponse "Token succesfully created"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /user/login [post]
func (h *Handler) login(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := h.Services.Login(ctx, req.Credentials, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrMismatchedPassword) || errors.Is(err, repository.ErrRecordNotFound):
			ctx.JSON(http.StatusUnauthorized, &api.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "user does not exists or password does not match",
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

	ctx.JSON(http.StatusCreated, &api.LoginResponse{
		Code:    http.StatusCreated,
		Message: "token succesfully created",
		Body:    token,
	})
}

// @Summary      Update user info
// @Tags         Users
// @Produce      json
// @Security ApiKeyAuth
// @Param data body api.LoginRequest true "Request body"
//
// @Success      201 {object} api.LoginResponse "User succesfully updated"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /user/update [patch]
func (h *Handler) updateUser(ctx *gin.Context) {
	var req api.UpdateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(int64)

	err = h.Services.UpdateUser(ctx, &entity.User{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password: entity.Password{
			Plaintext: req.Password,
		},
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
		Message: "user succesfully updated",
	})
}

// @Summary      Delete user
// @Tags         Users
// @Produce      json
// @Security ApiKeyAuth
//
// @Success      200 {object} api.DefaultResponse "User succesfully deleted"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /user/delete [delete]
func (h *Handler) deleteUser(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(int64)

	err := h.Services.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "user succesfully deleted",
	})
}
