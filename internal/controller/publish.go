package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Publish(c *gin.Context) {
	// token := c.Query("token")
	c.JSON(http.StatusOK, service.Response{
		StatusCode: 0,
	})
}

func (ctl *Controller) PublishList(c *gin.Context) {
	var rsp service.VideoListResponse
	token := c.PostForm("token")
	uc, err := jwtx.ParseUserClaims(token)
	if err != nil {
		ctl.logger.Sugar().Errorf("ParseUserClaims error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	r, err := ctl.service.PublicList(context.Background(), uc.UserID)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	rsp.Response = r.Response
	rsp.VideoList = r.VideoList
	c.JSON(http.StatusOK, rsp)
}
