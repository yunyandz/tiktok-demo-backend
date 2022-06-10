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
	apiRouter.GET("/feed/", middleware.JWTAuth(logger, false), ctl.Feed)

	// using jwt auth
	apiRouter.GET("/user/", middleware.JWTAuth(logger, false), ctl.UserInfo)

	apiRouter.POST("/user/register/", ctl.Register)
	apiRouter.POST("/user/login/", ctl.Login)

	apiRouter.POST("/publish/action/", middleware.JWTAuth(logger, true), ctl.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTAuth(logger, false), ctl.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JWTAuth(logger, true), ctl.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JWTAuth(logger, false), ctl.FavoriteList)

	apiRouter.POST("/comment/action/", middleware.JWTAuth(logger, true), ctl.CommentAction)
	apiRouter.GET("/comment/list/", middleware.JWTAuth(logger, false), ctl.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JWTAuth(logger, true), ctl.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JWTAuth(logger, false), ctl.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.JWTAuth(logger, false), ctl.FollowerList)
}
