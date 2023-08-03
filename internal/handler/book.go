package handler

import (
	"errors"
	"net/http"
	"one-lab-final/internal/entity"
	"one-lab-final/internal/handler/api"
	"one-lab-final/pkg/util"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

var (
	ErrEmptyTags = errors.New("tags must be present")
)

// @Summary      Create new book. Requires MODERATOR role or higher
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param data body api.CreateBookRequest true "Request body"
//
// @Success      201 {object} api.DefaultResponse "Book succesfully created"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/new [post]
func (h *Handler) createBook(ctx *gin.Context) {
	var req api.CreateBookRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.Tags = slices.DeleteFunc(req.Tags, func(tag string) bool {
		return tag == ""
	})
	if len(req.Tags) == 0 {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: ErrEmptyTags.Error(),
		})
		return
	}

	err = h.Services.CreateBook(ctx, &entity.Book{
		Title:       &req.Title,
		Description: &req.Description,
		Author:      &req.Author,
		Tags:        &req.Tags,
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
		Message: "book succesfully created",
	})
}

// @Summary      Get book by id
// @Tags         Books
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Book ID"
//
// @Success      200 {object} api.GetBookByIDResponse "Book info"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/{id} [get]
func (h *Handler) getBookByID(ctx *gin.Context) {
	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	book, err := h.Services.GetBookByID(ctx, id.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.GetBookByIDResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    book,
	})
}

// @Summary      Get books using pagination
// @Tags         Books
// @Produce      json
// @Param filter  query api.Filter true "Pagination filter"
//
// @Success      200 {object} api.GetBooksResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books [get]
func (h *Handler) getBooks(ctx *gin.Context) {
	var req api.GetBooksRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if req.Tags != nil {
		tags := strings.Split((*req.Tags)[0], ",")
		req.Tags = &tags
	}

	books, meta, err := h.Services.GetBooks(ctx,
		req.Title,
		req.Author,
		req.Tags,
		util.NewFilter(req.Page, req.PageSize, req.Sort),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.GetBooksResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    books,
		Meta:    *meta,
	})
}

// @Summary      Update book by id. Requires MODERATOR role or higher
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Book ID"
// @Param data body api.UpdateBookRequest true "Request body"
//
// @Success      200 {object} api.DefaultResponse "Book succesfully updated"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/update/{id} [patch]
func (h *Handler) updateBook(ctx *gin.Context) {
	var req api.UpdateBookRequest
	var id api.ID

	err := ctx.ShouldBindUri(&id)
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
	if req.Tags != nil {
		tags := slices.DeleteFunc(*req.Tags, func(tag string) bool {
			return tag == ""
		})
		req.Tags = &tags
		if len(*req.Tags) == 0 {
			ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: ErrEmptyTags.Error(),
			})
			return
		}

	}

	err = h.Services.UpdateBook(ctx, &entity.Book{
		ID:          id.Value,
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		Tags:        req.Tags,
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
		Message: "book succesfully updated",
	})

}

// @Summary      Delete book by ID. Requires MODERATOR role or higher
// @Tags         Books
// @Produce      json
// @Security ApiKeyAuth
// @Param        id   path      int  true  "Book ID"
//
// @Success      200 {object} api.DefaultResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/delete/{id} [delete]
func (h *Handler) deleteBook(ctx *gin.Context) {
	var id api.ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.Services.DeleteBook(ctx, id.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "book succesfully deleted",
	})
}
