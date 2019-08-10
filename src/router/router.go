package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SetUp(router *gin.Engine)  {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.GET("/health", func(context *gin.Context) {
		context.String(200,"hello AMS")
	})

	v1Account := router.Group("/account/v1.0")
	Account(v1Account)
	v1Message := router.Group("/message/v1.0")
	Message(v1Message)
}
