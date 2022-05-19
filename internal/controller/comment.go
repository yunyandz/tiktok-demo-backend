package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
)

func (ctl *Controller) CommentAction(c *gin.Context) {
	// token := c.Query("token")
	// sddddd

	c.JSON(http.StatusOK, service.Response{StatusCode: 0})
}

func (ctl *Controller) CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, service.CommentListResponse{
		Response: service.Response{StatusCode: 0},
	})
}
