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
	router.Use(middleware.CORSMiddleware())
	socket.StartSocket(router)
	v1Account := router.Group("/account/v1.0",middleware.Authrization())
	Account(v1Account)
	v1Message := router.Group("/message/v1.0",middleware.Authrization())
	Message(v1Message)
	v1Cloud := router.Group("/cloud/v1.0",middleware.Authrization())
	CloudApi(v1Cloud)
}
