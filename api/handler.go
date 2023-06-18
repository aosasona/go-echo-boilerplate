package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopi/api/auth"
	"gopi/pkg/response"
	"gopi/services"
)

type Handler struct {
	app     *echo.Echo
	service *services.Service
}

func New(app *echo.Echo, serviceUmbrella *services.Service) *Handler {
	return &Handler{
		app:     app,
		service: serviceUmbrella,
	}
}

func (h *Handler) MountRoutes() {
	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	auth.NewHandler(v1, h.service.NewAuthService()).Mount()

	v1.GET("/health", healthHandler)
	h.app.GET("/health", healthHandler)

	h.app.GET("*", func(c echo.Context) error {
		res := response.New(c)
		return res.Error(response.Data{Code: http.StatusNotFound, Message: "Not Found"})
	})
}

func healthHandler(c echo.Context) error {
	res := response.New(c)
	return res.Success(response.Data{Message: "I'm alive"})
}
