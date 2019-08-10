package router

import (
	"AMS/src/controllers"
	"github.com/gin-gonic/gin"
)

func Account(v1Router *gin.RouterGroup)  {

	v1Router.GET("/health", func(context *gin.Context) {
		context.String(200,"health account")
	})
	// 注册
	v1Router.POST("/logup",controllers.Logup)
	// 验证码检查手机号是否注册
	v1Router.POST("/checkphone",controllers.CheckPhone)
	// 验证码获取
	v1Router.GET("/code",controllers.GetCode)
	// 验证码登录
	v1Router.POST("/login/code",controllers.LoginWithCode)

}
