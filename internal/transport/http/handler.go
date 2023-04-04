package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"v001_onelab/internal/service" //nolint:typecheck

	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Handler struct {
	router      *echo.Echo
	UserService service.IUser
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UserService: service.User,
	}
}

func (h Handler) InitRouter() *echo.Echo {
	h.router = echo.New()

	h.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	api := h.router.Group("/api")
	users := api.Group("/users")

	users.GET("/", h.GetUsers)
	users.GET("/:id", h.GetUser)
	users.POST("/", h.CreateUser)
	users.PATCH("/:id", h.UpdateUser)
	users.DELETE("/:id", h.DeleteUser)

	return h.router
}
