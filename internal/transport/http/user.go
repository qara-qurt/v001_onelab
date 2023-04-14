package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"v001_onelab/internal/model"

	_ "v001_onelab/docs"
)

// GetUsers @Summary GetUsers
// @Tags users
// @Security ApiKeyAuth
// @Description get all users
// @ID get-users
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/ [get]
func (h *Handler) GetUsers(c echo.Context) error {
	res, err := h.service.User.GetAll()

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

// GetUser @Summary GetUser
// @Tags users
// @Security ApiKeyAuth
// @Description get user by ID
// @ID get-user
// @Accept json
// @Produce json
// @Param  id path int  true "Account ID"
// @Success 200 {object} model.UserResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/{id} [get]
func (h *Handler) GetUser(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid user ID"))
	}

	res, err := h.service.User.GetByID(id)
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

// SignUp @Summary SignUp
// @Tags auth
// @Description register
// @ID sign-up
// @Accept json
// @Produce json
// @Param  input body model.UserInput  true "User info"
// @Success 201
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c echo.Context) error {
	var user model.UserInput
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	err := h.service.User.Create(user)
	if err != nil {
		c.Logger().Error(err)
		status := http.StatusInternalServerError
		if errors.Is(err, model.ErrorAlreadyExist) {
			status = http.StatusBadRequest
		}
		return c.JSON(status, model.NewErrorResponse(err.Error()))
	}
	return c.NoContent(http.StatusCreated)
}

// SignIn @Summary SignIn
// @Tags auth
// @Description SignIn
// @ID sign-in
// @Accept json
// @Produce json
// @Param  input body model.SignInInput  true "User info"
// @Success 200 {object} model.Token
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c echo.Context) error {
	var user model.SignInInput
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	token, err := h.service.User.SignIn(user)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.Token{
		Token: token,
	})
}

// DeleteUser @Summary DeleteUser
// @Tags users
// @Security ApiKeyAuth
// @Description delete user
// @ID delete-user
// @Accept json
// @Produce json
// @Param  id path int  true "Account ID"
// @Success 200
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid user ID"))
	}

	if err := h.service.User.Delete(id); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

// UpdateUser @Summary UpdateUser
// @Tags users
// @Security ApiKeyAuth
// @Description update user
// @ID update-user
// @Accept json
// @Produce json
// @Param  input body model.UpdateUser  true "User info"
// @Success 200
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/ [patch]
func (h Handler) UpdateUser(c echo.Context) error {
	id := c.Get("id")

	var res model.UpdateUser
	if err := c.Bind(&res); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	if err := res.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	user := model.UserResponse{
		ID:       id.(uint),
		FullName: res.FullName,
		Login:    res.Login,
	}

	err := h.service.User.Update(user)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
	}
	return c.NoContent(http.StatusOK)
}

// ChangePassword @Summary ChangePassword
// @Tags users
// @Security ApiKeyAuth
// @Description update user
// @ID change-password
// @Accept json
// @Produce json
// @Param  input body model.ChangePassword  true "info"
// @Success 200
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/change-password [patch]
func (h Handler) ChangePassword(c echo.Context) error {
	var user model.ChangePassword
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	if err := user.Validate(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}

	if err := h.service.User.ChangePassword(user); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, model.ErrorPassword) {
			status = http.StatusBadRequest
		}
		return c.JSON(status, model.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
