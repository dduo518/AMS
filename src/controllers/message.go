package controllers

import (
	"AMS/src/app"
	"AMS/src/services"
	"github.com/gin-gonic/gin"
)

func SendVerifierMessage(c *gin.Context)  {
	phone:=c.PostForm("phone")
	result,err:= services.SendVerifierMessage(phone,"1234")
	if err!=nil{
		app.Response(c,err,100)
		return
	}
	app.Response(c,result,0)
}

