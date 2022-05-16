package service

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	db       *gorm.DB
	rds      *redis.Client
	logger   *zap.Logger
	producer sarama.AsyncProducer
}

var (
	service *Service
	once    sync.Once
)

// 启动一个新的service实例，当然是单例模式
func New(db *gorm.DB, rds *redis.Client, logger *zap.Logger, producer sarama.AsyncProducer) *Service {
	once.Do(func() {
		service = &Service{
			db:       db,
			rds:      rds,
			logger:   logger,
			producer: producer,
		}
	})
	return service
}
