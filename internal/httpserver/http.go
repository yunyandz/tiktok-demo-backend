package httpserver

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/router"
	"go.uber.org/zap"
)

func Run(config *config.Config, controller *controller.Controller, logger *zap.Logger) {
	r := gin.Default()

	router.InitRouter(r, controller)

	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
