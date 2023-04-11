package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_ "v001_onelab/docs"
	"v001_onelab/internal/model"
)

// CreateBook @Summary CreateBook
// @Tags books
// @Security ApiKeyAuth
// @Description create book
// @ID create-book
// @Accept json
// @Produce json
// @Param  input body model.BookInput  true "Book info"
// @Success 201
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /books/ [post]
func (h *Handler) CreateBook(c echo.Context) error {
	var book model.BookInput
	if err := c.Bind(&book); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	if err := book.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	if err := h.BookService.Create(book); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusCreated)

}

// GetBooks @Summary GetBooks
// @Tags books
// @Security ApiKeyAuth
// @Description get all books
// @ID get-books
// @Accept json
// @Produce json
// @Success 200 {object} []model.Book
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /books/ [get]
func (h *Handler) GetBooks(c echo.Context) error {
	books, err := h.BookService.GetAll()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, books)
}
