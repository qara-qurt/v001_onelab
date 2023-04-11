package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_ "v001_onelab/docs"
	"v001_onelab/internal/model"
)

// GetOrderBooks @Summary GetOrderBooks
// @Tags books
// @Security ApiKeyAuth
// @Description get order history books
// @ID get_order-books
// @Accept json
// @Produce json
// @Success 200 {object} []model.OrderBook
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /books/order-book [get]
func (h *Handler) GetOrderBooks(c echo.Context) error {
	books, err := h.OrderBookService.GetOrderBooks()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, books)

}

// GetOrderUserBooks @Summary ChangePassword
// @Tags user-books
// @Security ApiKeyAuth
// @Description get user books current
// @ID get-user-books
// @Accept json
// @Produce json
// @Success 200 {object} model.UserOrderBooksResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/books/current [get]
func (h Handler) GetOrderUserBooks(c echo.Context) error {
	userBooks, err := h.OrderBookService.GetOrderUserBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, userBooks)
}

// GetOrderUserBooksLastMounth @Summary ChangePassword
// @Tags user-books
// @Security ApiKeyAuth
// @Description get user books current
// @ID get-user-books-mounth
// @Accept json
// @Produce json
// @Success 200 {object} model.UserOrderBooksResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/books/last-mounth [get]
func (h Handler) GetOrderUserBooksLastMounth(c echo.Context) error {
	userBooks, err := h.OrderBookService.GetOrderUserBooksLastMounth()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, userBooks)
}
