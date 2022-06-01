package main

import (
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

// 这不用依赖注入框架真的好吗。。。
func main() {
	cfg, err := config.Phase()
	if err != nil {
		panic(err)
	}
	mylogger := logger.New(cfg)
	db := mysql.New(cfg, mylogger)
	rds := redis.New(cfg, mylogger)
	pdc := kafka.NewProducer(cfg)
	s3 := s3.New(cfg, mylogger)
	ser := service.New(cfg, db, rds, mylogger, pdc, s3)
	ctl := controller.New(ser, mylogger)
	httpserver.Run(cfg, ctl, mylogger)
}
