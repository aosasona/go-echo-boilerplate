package core

import (
	"gopi/internal/config"
	"gopi/internal/logger"

	"github.com/labstack/echo/v4"

	"github.com/uptrace/bun"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type Core struct {
	app    *echo.Echo
	db     *bun.DB
	cache  *bolt.DB
	config *config.Config
	logger *zap.Logger
}

func New() (*Core, error) {
	var err error
	c := new(Core)
	c.logger = logger.Logger
	return c, err
}

func (core *Core) SetDB(db *bun.DB) *Core {
	core.db = db
	return core
}

func (core *Core) SetCache(cache *bolt.DB) *Core {
	core.cache = cache
	return core
}

func (core *Core) SetConfig(config *config.Config) *Core {
	core.config = config
	return core
}
