package router

import (
	"github.com/gin-gonic/gin"
	"task05/controller"
	"task05/util"
)

func SetRouter() *gin.Engine {

	engine := gin.Default()

	{
		UserGroup := engine.Group("user")
		UserGroup.POST("register", controller.HandleUserRegister())
		UserGroup.POST("login", controller.HandleUserLogin())
	}

	{
		PostGroup := engine.Group("post")
		PostGroup.POST("add", util.JWTAuth(), controller.AddPostHandler())
		PostGroup.GET("all", controller.GetAllPost())
		PostGroup.GET(":id", controller.GetPostById())
		PostGroup.DELETE(":postId", util.JWTAuth(), controller.DeletePost())
		PostGroup.PUT("", util.JWTAuth(), controller.UpdatePost())
	}

	{
		CommentGroup := engine.Group("comment")
		CommentGroup.GET("get/:postId", controller.GetCommentsByPostId())
		CommentGroup.POST("add", util.JWTAuth(), controller.CreateCommentHandler())
	}

	return engine
}
