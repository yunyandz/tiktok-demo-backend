package redis

import (
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
)

var (
	rdb  *redis.Client
	once sync.Once
)

// 使用单例模式防止重复创建
func New(cfg *config.Config) *redis.Client {
	host := strings.Join([]string{cfg.Redis.Host, cfg.Redis.Port}, ":")
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: cfg.Redis.Password, // no password set
			DB:       cfg.Redis.DB,       // use default DB
		})
	})
	return rdb
}
