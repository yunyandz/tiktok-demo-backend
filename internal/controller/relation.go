package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

type RealationRequest struct {
	UserId     uint64 `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   uint64 `form:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

func (ctl *Controller) RelationAction(c *gin.Context) {
	// token := c.Query("token")
	var req RealationRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	switch req.ActionType {
	case 1:
		rsp := ctl.service.Follow(req.UserId, req.ToUserId)
		c.JSON(http.StatusOK, rsp)
	case 2:
		rsp := ctl.service.UnFollow(req.UserId, req.ToUserId)
		c.JSON(http.StatusOK, rsp)
	default:
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "invaild action"})
	}
	// c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

func (ctl *Controller) FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, service.UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}

func (ctl *Controller) FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, service.UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
	})
}
