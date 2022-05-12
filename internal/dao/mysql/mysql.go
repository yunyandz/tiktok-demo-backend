package mysql

import (
	"errors"
	"fmt"
	"sync"

	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

var (
	ErrCouldNotConnect = errors.New("could not connect to mysql")
)

// 使用单例模式防止重复创建
func New(cfg *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.Name, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
		var err error
		mysqlconfig := mysql.Config{
			DSN: dsn,
		}
		db, err = gorm.Open(mysql.New(mysqlconfig), &gorm.Config{})
		if err != nil {
			panic(ErrCouldNotConnect)
		}
	})
	return db
}
