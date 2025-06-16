package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task05/model"
	"task05/response"
)

func AddPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		post := new(model.Post)
		err := ctx.BindJSON(post)
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		}
		post.UserId = uint(ctx.MustGet("userID").(float64))
		err = model.CreatePost(post)
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		}
		response.Success(ctx, post)
	}
}

func GetAllPost() gin.HandlerFunc {
	return func(context *gin.Context) {
		posts, err := model.ListPosts()
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		response.Success(context, posts)
	}
}

func GetPostById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		parseUint, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
		post, err := model.GetPostById(uint(parseUint))
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		}
		response.Success(ctx, post)
	}
}

func DeletePost() gin.HandlerFunc {
	return func(context *gin.Context) {
		postId, err := strconv.ParseUint(context.Param("postId"), 10, 64)
		if err != nil {
			response.FailureWithMsg(context, "无效的postId")
			return
		}

		userID := uint(context.MustGet("userID").(float64))
		dbPost, err := model.GetPostById(uint(postId))
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		if dbPost.UserId != userID {
			response.FailureWithMsg(context, "无权删除该文章")
			return
		}
		model.DeletePost(uint(postId))
		response.Success(context, nil)
	}
}

func UpdatePost() gin.HandlerFunc {
	return func(context *gin.Context) {
		post := new(model.Post)
		err := context.BindJSON(post)
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		userID := uint(context.MustGet("userID").(float64))
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		dbPost, err := model.GetPostById(post.ID)
		if dbPost.UserId != userID {
			response.FailureWithMsg(context, "无权修改该文章")
			return
		}
		err = model.UpdatePost(post)
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		response.Success(context, nil)
	}

}
