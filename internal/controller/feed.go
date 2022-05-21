package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

// Feed same demo video list for every request
func (ctl *Controller) Feed(c *gin.Context) {
	lastestTime := time.Unix(c.PostForm("latest_time"))

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
	var rsp service.FeedResponse
	r, err := ctl.service.getFeed(context.Background(), uc.UserID, lastestTime)
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
