package controller

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
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

func (ctl *Controller) getUserClaims(c *gin.Context) (*jwtx.UserClaims, bool) {
	uc, e := c.Get("claims")
	if !e {
		ctl.logger.Sugar().Debugf("Get claims error: %v", e)
		return nil, false
	}
	return uc.(*jwtx.UserClaims), true
}
