package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task05/model"
	"task05/response"
)

func CreateCommentHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		comment := new(model.Comment)
		err := context.BindJSON(comment)
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		comment.UserID = uint(context.MustGet("userID").(float64))
		err = model.CreateComment(comment)
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		response.Success(context, comment)

	}
}

func GetCommentsByPostId() gin.HandlerFunc {
	return func(context *gin.Context) {
		postId, err := strconv.ParseUint(context.Param("postId"), 10, 64)
		if err != nil {
			response.FailureWithMsg(context, "无效的postId")
			return
		}
		comments, err := model.GetAllCommentsByPostId(uint(postId))
		if err != nil {
			response.FailureWithMsg(context, err.Error())
			return
		}
		response.Success(context, comments)
	}
}
