package main

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/mysql"
	"github.com/yunyandz/tiktok-demo-backend/internal/dao/redis"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func main() {
	cfg, err := config.Phase()
	if err != nil {
		panic(err)
	}
	mylogger := logger.New(cfg)
	db := mysql.New(cfg, mylogger)
	rds := redis.New(cfg)
	ser := service.New(db, rds, mylogger)
	ctl := controller.New(ser, mylogger)
	httpserver.Run(cfg, ctl, mylogger)
}
