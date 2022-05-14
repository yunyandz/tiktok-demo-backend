package mysql

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

var (
	once sync.Once
	db   *gorm.DB
)

var (
	ErrCouldNotConnect = errors.New("could not connect to mysql")
)

// 使用单例模式防止重复创建
func New(cfg *config.Config, logger *zap.Logger) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
		var err error
		mysqlconfig := mysql.Config{
			DSN: dsn,
		}
		logger := zapgorm2.New(logger)
		logger.SetAsDefault()
		db, err = gorm.Open(mysql.New(mysqlconfig), &gorm.Config{Logger: logger})
		if err != nil {
			panic(ErrCouldNotConnect)
		}
		if err = db.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}); err != nil {
			panic(err)
		}
	})
	return db
}
