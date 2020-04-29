package routers

import (
	"github.com/gin-gonic/gin"

	"go-space/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(),middlewares.EnabledCORS())

	authGroup := r.Group("/auth")
	authGroup.Use(middlewares.AuthRequired())
	{
		authGroup.POST("/test", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"msg":"success",
			})
		})
		authGroup.GET("/user",GetUserInfo)
	}

	r.GET("/", func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.String(200,"<h1>Hello World! My name is Gin Powered By Golang!</h1>")
	})
	r.GET("/user", Add)
	r.GET("/get",GetAllCountries)
	r.POST("/register",RegisterUser)
	r.POST("/logon",LogonUser)
	r.GET("/logout",Logout)
	r.GET("/getMovies",GetMovies)
	// 微信公众号验证
	r.GET("/weChatTestSign",WeChatTestSign)
	// 前端登录地址
	r.GET("/weChatLogin",WeChatLogin)
	// 微信回调地址
	r.GET("/weChatGetToken",WeChatGetToken)

	r.GET("/xml", func(context *gin.Context) {
		context.XML(200,gin.H{
			"code": 200,
			"data": "1",
			"msg":"a",
		})
	})
	return r
}