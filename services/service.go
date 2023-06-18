package services

import (
	"gopi/internal/config"

	"github.com/uptrace/bun"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type Service struct {
	db     *bun.DB
	cache  *bolt.DB
	config *config.Config
	logger *zap.Logger
}

type BaseServiceInterface interface {
	Config() *config.Config
	Logger() *zap.Logger
}

func New(db *bun.DB, cache *bolt.DB, config *config.Config, logger *zap.Logger) *Service {
	return &Service{
		db:     db,
		cache:  cache,
		config: config,
		logger: logger,
	}
}

func (s *Service) Config() *config.Config {
	return s.config
}

func (s *Service) Logger() *zap.Logger {
	return s.logger
}

func (s *Service) NewAuthService() *authService {
	return &authService{s}
}
