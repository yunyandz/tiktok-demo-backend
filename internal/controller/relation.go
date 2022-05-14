package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) RelationAction(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

func (ctl *Controller) FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}

func (ctl *Controller) FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
