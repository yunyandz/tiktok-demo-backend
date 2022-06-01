package logger

import (
	"sync"

	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	suger  *zap.SugaredLogger
	once   sync.Once
)

// 创建新的logger
func New(cfg *config.Config) *zap.Logger {
	once.Do(func() {
		if !cfg.Debug {
			logger, _ = zap.NewProduction()
		} else {
			logger, _ = zap.NewDevelopment()
			logger.Debug("running in debug mode...")
		}
		suger = logger.Sugar()
	})
	return logger
}

// 返回一个suger模式下的logger，可以引入后直接使用Suger().xxx()
func Suger() *zap.SugaredLogger {
	return suger
}
