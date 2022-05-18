package httpserver

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver/router"

	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
)

func Run(config *config.Config, controller *controller.Controller, logger *zap.Logger) {
	r := gin.Default()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	router.InitRouter(r, controller)

	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
