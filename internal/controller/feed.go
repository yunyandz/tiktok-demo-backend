package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

// Feed same demo video list for every request
func (ctl *Controller) Feed(c *gin.Context) {
	c.JSON(http.StatusOK, service.FeedResponse{
		Response: service.Response{StatusCode: 0},
		NextTime: time.Now().Unix(),
	})
}
