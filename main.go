package main

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/mysql"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/redis"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
)

func main() {
	cfg, err := config.Phase()
	if err != nil {
		panic(err)
	}
	logger.New(cfg)
	mysql.New(cfg)
	redis.New(cfg)
	httpserver.Run(cfg)
}
