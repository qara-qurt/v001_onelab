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

func New(service *service.Service) *Handler {
	return &Handler{
		router:      echo.New(),
		UserService: service.User,
	}
}

func (h Handler) InitRouter() *echo.Echo {
	h.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := h.router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	users := api.Group("/users")
	{
		users.Use(h.authMiddleware)
		users.GET("/", h.GetUsers)
		users.GET("/:id", h.GetUser)
		users.PATCH("/change-password", h.ChangePassword)
		users.PATCH("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
	return h.router
}
