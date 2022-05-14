package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) FavoriteAction(c *gin.Context) {
	// token := c.Query("token")

	c.JSON(http.StatusOK, service.Response{StatusCode: 0})
}

func (ctl *Controller) FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, service.VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
