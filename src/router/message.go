package router

import (
	"AMS/src/controllers"
	"github.com/gin-gonic/gin"
)

func Message(router *gin.RouterGroup)  {

	router.POST("/send",controllers.SendVerifierMessage)
}