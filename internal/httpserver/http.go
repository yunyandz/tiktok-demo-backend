package httpserver

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver/router"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
)

func Run(config *config.Config, controller *controller.Controller, logger *zap.Logger) {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	if config.Debug {
		pprof.Register(r)
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	router.InitRouter(logger, r, controller)

	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
