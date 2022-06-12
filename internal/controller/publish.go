package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Publish(c *gin.Context) {
	uc, _ := ctl.getUserClaims(c)
	var err error
	if err != nil {
		c.JSON(http.StatusBadRequest, service.Response{
			StatusCode: -1,
			StatusMsg:  errorx.ErrReadVideo.Error(),
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		ctl.logger.Sugar().Errorf("FormFile error: %v", err)
		c.JSON(http.StatusBadRequest, service.Response{
			StatusCode: -1,
			StatusMsg:  errorx.ErrReadVideo.Error(),
		})
		return
	}
	file, err := data.Open()
	if err != nil {
		ctl.logger.Sugar().Errorf("data.Open error: %v", err)
		c.JSON(http.StatusBadRequest, service.Response{
			StatusCode: -1,
			StatusMsg:  errorx.ErrReadVideo.Error(),
		})
		return
	}
	title := c.PostForm("title")
	if data.Header.Get("Content-Type") != "video/mp4" {
		ctl.logger.Sugar().Errorf("data.Header.Get(Content-Type) error: %v", err)
		c.JSON(http.StatusBadRequest, service.Response{
			StatusCode: -1,
			StatusMsg:  errorx.ErrReadVideo.Error(),
		})
		return
	}
	res := ctl.service.PublishVideo(context.Background(), uc.UserID, data.Filename, file, title)
	c.JSON(http.StatusOK, service.Response{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	})
}

type PublishListRequest struct {
	UserID uint64 `form:"user_id"`
}

func (ctl *Controller) PublishList(c *gin.Context) {
	var rsp service.VideoListResponse
	selfId := uint64(0)
	uc, e := ctl.getUserClaims(c)
	if e {
		selfId = uc.UserID
	}
	var req PublishListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ctl.logger.Sugar().Errorf("ShouldBindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, service.Response{
			StatusCode: -1,
			StatusMsg:  errorx.ErrReadVideo.Error(),
		})
		return
	}
	r := ctl.service.GetVideoList(context.Background(), selfId, req.UserID)
	rsp.Response = r.Response
	rsp.VideoList = r.VideoList
	c.JSON(http.StatusOK, rsp)
}
