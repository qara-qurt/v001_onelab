package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"v001_onelab/internal/service" //nolint:typecheck

	_ "v001_onelab/docs"
)

type Handler struct {
	router  *echo.Echo
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		router:  echo.New(),
		service: service,
	}
}

func (h Handler) InitRouter() *echo.Echo {
	h.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := h.router.Group("/api")

	api.GET("/swagger/*", echoSwagger.WrapHandler)
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
		users.PATCH("/", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)

		userBooks := users.Group("/books")
		{
			userBooks.GET("/current", h.GetOrderUserBooks)
			userBooks.GET("/last-mounth", h.GetOrderUserBooksLastMounth)
		}
	}

	books := api.Group("/books")
	{
		books.Use(h.authMiddleware)
		books.POST("/", h.CreateBook)
		books.GET("/", h.GetBooks)
		//Общая истроия выдача книг
		books.GET("/order-book", h.GetOrderBooks)
	}
	return h.router
}
