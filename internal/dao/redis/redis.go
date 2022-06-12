package redis

import (
	"context"
	"strings"
	"sync"
	"time"

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
		var err error
		for i := 0; i < 3; i++ {
			if _, err = rdb.Ping(context.TODO()).Result(); err == nil {
				break
			}
			logger.Error("connect to redis failed", zap.Error(err), zap.Int("retry", i))
			time.Sleep(time.Second * 3)
		}
		if err != nil {
			panic(err)
		}
		logger.Info("connect to redis success")
		logger.Sugar().Infof("redis connect success, host: %s, database: %d", host, cfg.Redis.Database)
	})
	return rdb
}
