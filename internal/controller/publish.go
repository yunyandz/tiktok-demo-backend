package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Publish(c *gin.Context) {
	token := c.PostForm("token")
	uc, err := jwtx.ParseUserClaims(token)
	if err != nil {
		ctl.logger.Sugar().Errorf("ParseUserClaims error: %v", err)
		c.JSON(http.StatusUnauthorized, service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
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
	// TODO: 判断文件类型是否为视频文件
	res := ctl.service.PublishVideo(context.Background(), uc.UserID, data.Filename, file)
	c.JSON(http.StatusOK, service.Response{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	})
}

func (ctl *Controller) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, service.VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
