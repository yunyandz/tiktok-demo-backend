package controller

import (
	"sync"

	"github.com/yunyandz/tiktok-demo-backend/internal/service"
	"go.uber.org/zap"
)

var (
	controller *Controller
	once       sync.Once
)

type Controller struct {
	service *service.Service
	logger  *zap.Logger
}

func New(service *service.Service, logger *zap.Logger) *Controller {
	once.Do(func() {
		controller = &Controller{
			service: service,
			logger:  logger,
		}
	})
	return controller
}
