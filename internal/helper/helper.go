package helper

import (
	"gopi/internal/config"

	"go.uber.org/zap"
)

type helper struct {
	config *config.Config
	logger *zap.Logger
}

func New(config *config.Config, logger *zap.Logger) *helper {
	return &helper{
		config: config,
		logger: logger,
	}
}
