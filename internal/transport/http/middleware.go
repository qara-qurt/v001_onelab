package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromRequest(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		userId, err := h.UserService.ParseToken(c.Request().Context(), token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		c.Set("id", userId)
		return next(c)
	}
}

func getTokenFromRequest(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}
