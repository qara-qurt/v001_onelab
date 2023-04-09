package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"v001_onelab/internal/model"
)

func (h *Handler) CreateBook(c echo.Context) error {
	var book model.BookInput
	if err := c.Bind(&book); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := book.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := h.BookService.Create(book); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "book created",
	})

}

func (h *Handler) GetBooks(c echo.Context) error {
	books, err := h.BookService.GetAll()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, books)
}

func (h *Handler) GetOrderBooks(c echo.Context) error {
	books, err := h.BookService.GetOrderBooks()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, books)

}
