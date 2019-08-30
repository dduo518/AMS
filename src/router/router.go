package router

import (
	"AMS/src/middleware"
	"AMS/src/socket"
	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine)  {

	router.GET("/health", func(context *gin.Context) {
		context.String(200,"hello AMS 222")
	})

	socket.StartSocket(router)

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.Authrization())

	v1Account := router.Group("/account/v1.0")
	Account(v1Account)
	v1Message := router.Group("/message/v1.0")
	Message(v1Message)
	v1Cloud := router.Group("/cloud/v1.0")
	CloudApi(v1Cloud)
}
