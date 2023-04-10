package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"v001_onelab/internal/model"
)

func (h *Handler) GetUsers(c echo.Context) error {
	res, err := h.UserService.GetAll()

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetUser(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid user ID",
		})
	}

	res, err := h.UserService.GetByID(id)
	if err != nil {
		c.Logger().Error(err)
		status := http.StatusInternalServerError
		if err == model.ErrorNotFound {
			status = http.StatusBadRequest
		}
		return c.JSON(status, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) SignUp(c echo.Context) error {
	var user model.UserInput
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := h.UserService.Create(user)
	if err != nil {
		c.Logger().Error(err)
		status := http.StatusInternalServerError
		if errors.Is(err, model.ErrorAlreadyExist) {
			status = http.StatusBadRequest
		}
		return c.JSON(status, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "user was created",
	})
}

func (h *Handler) SignIn(c echo.Context) error {
	var user model.SignInInput
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	token, err := h.UserService.SignIn(user)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) DeleteUser(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid user ID",
		})
	}

	if err := h.UserService.Delete(id); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted",
	})
}

func (h Handler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid user ID",
		})
	}

	var user model.UserResponse
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	user.ID = uint(id)
	err = h.UserService.Update(user)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "user updated",
	})
}

func (h Handler) ChangePassword(c echo.Context) error {
	var user model.ChangePassword
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := h.UserService.ChangePassword(user); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, model.ErrorPassword) {
			status = http.StatusBadRequest
		}
		return c.JSON(status, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "password changed",
	})
}

func (h Handler) GetOrderUserBooks(c echo.Context) error {
	userBooks, err := h.UserService.GetOrderUserBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userBooks)
}

func (h Handler) GetOrderUserBooksLastMounth(c echo.Context) error {
	userBooks, err := h.UserService.GetOrderUserBooksLastMounth()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userBooks)
}
