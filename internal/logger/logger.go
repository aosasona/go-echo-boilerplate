package logger

import (
	"gopi/internal/config"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error
	c, err := config.LoadEnv(".")

	if c.AppEnv == config.DEVELOPMENT {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}

	if err != nil {
		panic("failed to create logger")
	}
}
