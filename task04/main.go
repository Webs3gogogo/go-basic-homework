package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	// 简单的路由组: v1
	{
		v1 := router.Group("/v1", aroundHandler())
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	router.GET("/ping", func(c *gin.Context) {
		url := c.Request.URL
		c.JSON(200, gin.H{
			"message": "pong",
			"url":     url.String(),
		})
	})
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	router.GET("/xml", func(context *gin.Context) {
		context.XML(200, gin.H{
			"message": "This is an XML response",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func aroundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 这里可以添加一些中间件逻辑，比如日志记录、认证等
		// 例如打印请求方法和路径
		method := c.Request.Method
		path := c.Request.URL.Path
		fmt.Println("Request Method:", method, "Path:", path)
		// 在请求处理前执行的逻辑
		c.Next() // 调用下一个处理器

		// 在请求处理后执行的逻辑
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors.String()})
		}
	}

}

func loginEndpoint(c *gin.Context) {

	if checkBodyType(c) {
		return
	}

	// 登录逻辑
	c.JSON(200, gin.H{
		"message": "Login successful",
	})
}

func checkBodyType(ctx *gin.Context) bool {
	s1 := student{}
	if error := ctx.ShouldBindBodyWithJSON(&s1); error != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid JSON body",
		})
		return false
	}

	return true
}

func submitEndpoint(ctx *gin.Context) {
	// 提交逻辑
	ctx.JSON(200, gin.H{
		"message": "Submit successful",
	})
}

func readEndpoint(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Read successful",
	})
}
