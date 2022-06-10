package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

type RegisterRequest struct {
	Username string `form:"username" binding:"required,min=3,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

type RegisterResponse struct {
	service.Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (ctl *Controller) Register(c *gin.Context) {
	var req RegisterRequest
	var rsp RegisterResponse
	err := c.ShouldBindQuery(&req)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	r, err := ctl.service.Register(req.Username, req.Password)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	rsp.Response = r.Response
	rsp.UserID = r.UserID
	rsp.Token = r.Token
	c.JSON(http.StatusOK, rsp)
}

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	service.Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (ctl *Controller) Login(c *gin.Context) {
	var req LoginRequest
	var rsp LoginResponse
	err := c.ShouldBindQuery(&req)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	r, err := ctl.service.Login(req.Username, req.Password)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	rsp.Response = r.Response
	rsp.UserID = r.UserID
	rsp.Token = r.Token
	c.JSON(http.StatusOK, rsp)
}

type UserInfoRequest struct {
	UserID uint64 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserInfoResponse struct {
	service.Response
	service.User `json:"user"`
}

func (ctl *Controller) UserInfo(c *gin.Context) {
	var req UserInfoRequest
	var rsp UserInfoResponse
	err := c.ShouldBindQuery(&req)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	selfId := uint64(0)
	uc, e := ctl.getUserClaims(c)
	if e {
		selfId = uc.UserID
	}

	r, err := ctl.service.GetUserInfo(selfId, req.UserID)
	if err != nil {
		rsp.Response = service.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	rsp.Response = r.Response
	rsp.User = r.User
	c.JSON(http.StatusOK, rsp)

}
