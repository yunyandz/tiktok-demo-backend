package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
