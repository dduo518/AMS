package controllers

import (
	"AMS/src/_interface"
	"AMS/src/app"
	service "AMS/src/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 注册
func Logup(c *gin.Context)  {
	var logupParams _interface.Account
	if err:=c.ShouldBind(&logupParams);err!=nil{
		app.Response(c,err.Error(),400)
		return
	}
	result,err:=service.LogupServices(logupParams);
	if err!=nil{
		app.Response(c,err.Error(),400)
		return
	}
	logupParams.UserId = result
	app.Response(c,logupParams,0)
}

// 获取验证码登录 3min 有效
func GetCode(c *gin.Context) {
	phone:=c.Query("phone")
	fmt.Println(phone)
	result,err:=service.GetCodeWithLogin(phone);
	if err!=nil{
		app.Response(c,err,400)
		return
	}
	app.Response(c,result,0)
}

// 使用验证码登录
func LoginWithCode(c *gin.Context)  {
	phone:=c.Query("phone")
	code:=c.PostForm("code")
	account,err:=service.VerifyCode(phone,code,"appid")
	if err!=nil{
		app.Response(c,err,400);
		return
	}
	app.Response(c,account,0)
}

// 检查手机是否已经被注册
func CheckPhone(c *gin.Context){
	phone:=c.PostForm("phone");
	result:= service.CheckPhone(phone);
	app.Response(c,result,0);
}