package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gopi/api"
	"gopi/pkg/response"
	"gopi/services"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// InitApp ONLY CALL THIS WHEN OTHER DEPENDENCIES ARE SET
func (core *Core) InitApp() {
	app := echo.New()

	app.JSONSerializer = CustomJSONSerializer{}
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		core.logger.Error("Fatal error", zap.Error(err))
		_ = handleUnhandledError(c, err)
	}

	core.app = app
}

func (core *Core) Run() error {
	conf := core.config

	core.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{conf.AllowedOrigins},
	}))

	core.app.Use(middleware.Logger())
	core.app.Use(middleware.Recover())

	sv := services.New(core.db, core.cache, core.config, core.logger)
	handler := api.New(core.app, sv)
	handler.MountRoutes()

	log.Infof("Listening on port %v", conf.Port)
	return core.app.Start(fmt.Sprintf("0.0.0.0:%v", conf.Port))
}

func (core *Core) Kill() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := core.app.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func (core *Core) CloseConnections() error {
	if err := core.db.Close(); err != nil {
		return err
	}
	if err := core.cache.Close(); err != nil {
		return err
	}
	return nil
}

func handleUnhandledError(c echo.Context, err error) error {
	if err != nil {
		if e, ok := err.(*echo.HTTPError); ok && e.Code < 500 {
			if msg, ok := e.Message.(string); ok {
				return response.New(c).Error(response.Data{Code: e.Code, Error: msg})
			}
			return response.New(c).Error(response.Data{Code: e.Code, Error: "Bad request"})
		}

		return response.New(c).Error(response.Data{Code: http.StatusInternalServerError, Error: "Internal server error :("})
	}

	return nil
}
