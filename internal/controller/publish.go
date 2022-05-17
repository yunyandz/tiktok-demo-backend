package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Publish(c *gin.Context) {
	token := c.Query("token")
	uc, err := jwtx.ParseUserClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	file, err := data.Open()
	ctl.service.PublishVideo(uc.UserID, data.Filename, file)
	c.JSON(http.StatusOK, service.Response{
		StatusCode: 0,
	})
}

func (ctl *Controller) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, service.VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
