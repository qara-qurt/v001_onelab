package http

import (
	"github.com/julienschmidt/httprouter"
	"v001_onelab/internal/service" //nolint:typecheck
)

type Handler struct {
	router      *httprouter.Router
	UserService service.IUser
}

func New(service *service.Service) *Handler {
	return &Handler{
		UserService: service.User,
	}
}

func (h Handler) InitRouter() *httprouter.Router {
	h.router = httprouter.New()

	h.router.GET("/api/users/", h.loggingMiddleware(h.GetUsers))
	h.router.GET("/api/users/:id", h.loggingMiddleware(h.GetUser))
	h.router.POST("/api/users/", h.loggingMiddleware(h.CreateUser))
	h.router.PATCH("/api/users/:id", h.loggingMiddleware(h.UpdateUser))
	h.router.DELETE("/api/users/:id", h.loggingMiddleware(h.DeleteUser))

	return h.router
}
