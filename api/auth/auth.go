package auth

import (
	"github.com/labstack/echo/v4"
	"gopi/services"
)

type Service interface {
	services.BaseServiceInterface
	Signin() error
	Signup() error
}

type Handler struct {
	router  *echo.Group
	service Service
}

func NewHandler(router *echo.Group, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (authHandler *Handler) Mount() {
	router := authHandler.router.Group("/auth")
	router.POST("/sign-in", authHandler.signin)
	router.POST("/sign-up", authHandler.signup)
}

func (authHandler *Handler) signin(c echo.Context) error {
	return nil
}

func (authHandler *Handler) signup(c echo.Context) error {
	return nil
}
