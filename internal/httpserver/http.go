package httpserver

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/router"
)

func Run(config *config.Config) {
	r := gin.Default()

	router.InitRouter(r)

	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
