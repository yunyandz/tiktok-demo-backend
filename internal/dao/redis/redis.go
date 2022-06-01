package redis

import (
	"context"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"go.uber.org/zap"
)

var (
	rdb  *redis.Client
	once sync.Once
)

// 使用单例模式防止重复创建
func New(cfg *config.Config, logger *zap.Logger) *redis.Client {
	host := strings.Join([]string{cfg.Redis.Host, cfg.Redis.Port}, ":")
	once.Do(func() {
		if !cfg.Redis.Vaild {
			rdb = nil
			return
		}
		rdb = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.Database,
		})
		if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
			panic(err)
		}
		logger.Sugar().Infof("redis connect success, host: %s, database: %d", host, cfg.Redis.Database)
	})
	return rdb
}
