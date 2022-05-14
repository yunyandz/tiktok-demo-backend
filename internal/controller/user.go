package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) Register(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")

	// token := username + password

	c.JSON(http.StatusOK, service.UserLoginResponse{
		Response: service.Response{StatusCode: 0},
	})
}

func (ctl *Controller) Login(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")

	// token := username + password

	c.JSON(http.StatusOK, service.UserLoginResponse{
		Response: service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})
}

func (ctl *Controller) UserInfo(c *gin.Context) {
	// token := c.Query("token")

	c.JSON(http.StatusOK, service.UserResponse{
		Response: service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})
}
