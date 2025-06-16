package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"` // HTTP status code
	Message string      `json:"msg"`  // Response message
	Data    interface{} `json:"data"` // Response data, can be any type
}

type ResultCodeItem struct {
	Code    int
	Message string
}

var ResultCode = struct {
	SUCCESS ResultCodeItem
	FAILURE ResultCodeItem
}{
	SUCCESS: ResultCodeItem{Code: 200, Message: "Success"},
	FAILURE: ResultCodeItem{Code: 500, Message: "Failure"},
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Code:    ResultCode.SUCCESS.Code,
		Message: ResultCode.SUCCESS.Message,
		Data:    data,
	})
}

func SuccessWithMsg(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Code:    ResultCode.SUCCESS.Code,
		Message: msg,
		Data:    data,
	})
}

func Failure(ctx *gin.Context) {
	ctx.JSON(200, Response{
		Code:    ResultCode.FAILURE.Code,
		Message: ResultCode.FAILURE.Message,
		Data:    nil,
	})
}
func FailureWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(200, Response{
		Code:    ResultCode.FAILURE.Code,
		Message: msg,
		Data:    nil,
	})
}

func FailureWithResultCode(ctx *gin.Context, resultCode ResultCodeItem) {
	ctx.JSON(200, Response{
		Code:    resultCode.Code,
		Message: resultCode.Message,
		Data:    nil,
	})
}

func FailureWithMsgAndResultCode(ctx *gin.Context, msg string, resultCode ResultCodeItem) {
	ctx.JSON(200, Response{
		Code:    resultCode.Code,
		Message: msg,
		Data:    nil,
	})

}
