package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

type RealationRequest struct {
	Token      string `form:"token" binding:"required"`
	ToUserId   uint64 `form:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

const (
	Follow   = 1
	UnFollow = 2
)

func (ctl *Controller) RelationAction(c *gin.Context) {
	// token := c.Query("token")
	var req RealationRequest
	err := c.ShouldBind(&req)
	uc, _ := ctl.getUserClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if req.ToUserId == uc.UserID {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "不能对自己进行操作"})
		return
	}
	switch req.ActionType {
	case Follow:
		rsp := ctl.service.Follow(uc.UserID, req.ToUserId)
		c.JSON(http.StatusOK, rsp)
	case UnFollow:
		rsp := ctl.service.UnFollow(uc.UserID, req.ToUserId)
		c.JSON(http.StatusOK, rsp)
	default:
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "invaild action"})
	}
	// c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

type FollowRequest struct {
	UserId uint64 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func (ctl *Controller) FollowList(c *gin.Context) {
	var req FollowRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.UserListResponse{
			Response: service.Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	selfId := uint64(0)
	uc, e := ctl.getUserClaims(c)
	if e {
		selfId = uc.UserID
	}

	rsp := ctl.service.GetFollowList(selfId, req.UserId)
	c.JSON(http.StatusOK, rsp)
}

func (ctl *Controller) FollowerList(c *gin.Context) {
	var req FollowRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.UserListResponse{
			Response: service.Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	selfId := uint64(0)
	uc, e := ctl.getUserClaims(c)
	if e {
		selfId = uc.UserID
	}

	rsp := ctl.service.GetFollowerList(selfId, req.UserId)
	c.JSON(http.StatusOK, rsp)
}
