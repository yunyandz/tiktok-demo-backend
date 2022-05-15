package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
)

func InitRouter(r *gin.Engine, ctl *controller.Controller) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", ctl.Feed)
	apiRouter.GET("/user", ctl.UserInfo)
	apiRouter.POST("/user/register", ctl.Register)
	apiRouter.POST("/user/login", ctl.Login)
	apiRouter.POST("/publish/action/", ctl.Publish)
	apiRouter.GET("/publish/list/", ctl.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", ctl.FavoriteAction)
	apiRouter.GET("/favorite/list/", ctl.FavoriteList)
	apiRouter.POST("/comment/action/", ctl.CommentAction)
	apiRouter.GET("/comment/list/", ctl.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", ctl.RelationAction)
	apiRouter.GET("/relation/follow/list/", ctl.FollowList)
	apiRouter.GET("/relation/follower/list/", ctl.FollowerList)
}
