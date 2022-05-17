package service

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	s3Object "github.com/yunyandz/tiktok-demo-backend/internal/s3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	cfg      *config.Config
	db       *gorm.DB
	rds      *redis.Client
	logger   *zap.Logger
	producer sarama.AsyncProducer
	s3       s3Object.S3ObjectAPI
}

var (
	service *Service
	once    sync.Once
)

// 启动一个新的service实例，当然是单例模式
func New(cfg *config.Config, db *gorm.DB,
	rds *redis.Client,
	logger *zap.Logger,
	producer sarama.AsyncProducer,
	s3 s3Object.S3ObjectAPI) *Service {
	once.Do(func() {
		service = &Service{
			cfg:      cfg,
			db:       db,
			rds:      rds,
			logger:   logger,
			producer: producer,
			s3:       s3,
		}
	})
	return service
}
