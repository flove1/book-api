package handler

import (
	"net/http"
	"one-lab-final/internal/handler/api"

	"github.com/gin-gonic/gin"
)

// @Summary      Check if server is running
// @Tags         Healthcheck
// @Produce      json
//
// @Success      200 {object} api.DefaultResponse
// @Failure      494  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /healthcheck [get]
func (h *Handler) healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "ok",
	})
}
