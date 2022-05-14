package service

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.Logger
}

var (
	service *Service
	once    sync.Once
)

// 启动一个新的service实例，当然是单例模式
func New(db *gorm.DB, rds *redis.Client, logger *zap.Logger) *Service {
	once.Do(func() {
		service = &Service{
			db:     db,
			rds:    rds,
			logger: logger,
		}
	})
	return service
}
