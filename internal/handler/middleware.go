package handler

import (
	"errors"
	"net/http"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/handler/api"
	"one-lab-final/internal/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) requireRole(requiredRole entity.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok := h.authenticate(ctx)
		if !ok {
			return
		}

		value, _ := ctx.Get("role")

		role := value.(entity.Role)
		if role < requiredRole {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "denied",
			})
			return
		}

		ctx.Next()
	}
}

func (h *Handler) requireAuthenticatedUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok := h.authenticate(ctx)
		if !ok {
			return
		}

		ctx.Next()
	}
}

func (h *Handler) authenticate(ctx *gin.Context) bool {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "authentication is required",
		})
		return false
	}

	token := strings.Split(authHeader, " ")
	if len(token) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "authorization header is malformed",
		})
		return false
	}

	user, err := h.Services.GetUserByToken(ctx, token[1])
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			})
			return false
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return false
		}
	}

	ctx.Set("userID", user.ID)
	ctx.Set("role", user.Role)

	return true
}
