package controller

import (
	"github.com/gin-gonic/gin"
	"task05/model"
	"task05/response"
	"task05/util"
)

func HandleUserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		}
		err = model.Login(&user)
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		}
		token, _ := util.GenerateToken(user.ID, []string{"user"})
		response.Success(ctx, map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
			"token":    token,
		})
	}
}

func HandleUserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			response.Failure(ctx)
			return
		}
		err = model.RegisterUser(&user)
		if err != nil {
			response.FailureWithMsg(ctx, err.Error())
			return
		} else {
			token, err := util.GenerateToken(user.ID, []string{"user"})
			if err != nil {
				response.FailureWithMsg(ctx, err.Error())
				return
			}
			response.Success(ctx, token)
		}
	}

}
