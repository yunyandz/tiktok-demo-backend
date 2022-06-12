package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/mysql"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/redis"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver"
	"github.com/yunyandz/tiktok-demo-backend/internal/kafka"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
	"github.com/yunyandz/tiktok-demo-backend/internal/s3"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

// 用了 好用
func main() {
	app := fx.New(
		fx.Provide(
			config.Phase,
			logger.New,
			mysql.New,
			redis.New,
			kafka.NewProducer,
			s3.New,
			service.New,
			controller.New,
		),
		fx.Invoke(
			httpserver.Run,
		),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	)
	app.Run()
}
