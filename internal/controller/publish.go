package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Publish(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.Response{
		StatusCode: 0,
	})
}

func (ctl *Controller) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
