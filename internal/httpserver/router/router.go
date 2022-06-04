package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/controller"
	"github.com/yunyandz/tiktok-demo-backend/internal/httpserver/middleware"
	"go.uber.org/zap"
)

func InitRouter(logger *zap.Logger, r *gin.Engine, ctl *controller.Controller) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", ctl.Feed)

	// using jwt auth
	apiRouter.GET("/user", middleware.JWTAuth(logger), ctl.UserInfo)

	apiRouter.POST("/user/register", ctl.Register)
	apiRouter.POST("/user/login", ctl.Login)

	apiRouter.POST("/publish/action/", middleware.JWTAuth(logger), ctl.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTAuth(logger), ctl.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JWTAuth(logger), ctl.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JWTAuth(logger), ctl.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.JWTAuth(logger), ctl.CommentAction)
	apiRouter.GET("/comment/list/", ctl.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JWTAuth(logger), ctl.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JWTAuth(logger), ctl.FollowList)
	apiRouter.GET("/relation/follower/list/", ctl.FollowerList)
}
