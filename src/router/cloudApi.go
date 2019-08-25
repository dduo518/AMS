package router

import (
	"AMS/src/controllers"
	"github.com/gin-gonic/gin"
)
func CloudApi(v1Router *gin.RouterGroup){

	v1Router.GET("/credentials/health", func(context *gin.Context) {
		context.String(200,"health account")
	})

	v1Router.GET("/credentials", controllers.GetFederationToken)

	v1Router.POST("/upload", controllers.UploadFile)
}