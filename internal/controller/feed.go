package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

// Feed same demo video list for every request
func (ctl *Controller) Feed(c *gin.Context) {
	isnew := true
	lasttimestring := c.Query("lasttimestamp")
	lasttimestamp := 0
	var lastestTime time.Time
	var err error
	if lasttimestring != "" {
		lasttimestamp, err = strconv.Atoi(lasttimestring)
		if err != nil {
			ctl.logger.Sugar().Errorf("Atoi error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			return
		}
		lastestTime = time.Unix(int64(lasttimestamp), 0)
		isnew = false
	}

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
	r, err := ctl.service.GetFeed(context.Background(), uc.UserID, isnew, lastestTime)
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
