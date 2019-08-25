package middleware

import (
	"AMS/src/app"
	"AMS/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT，DELETE")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Authrization() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token ==""{
			token = c.Query("accessToken")
		}
		if token ==""{
			app.Response(c,"缺少签名",400)
			c.Abort()
			return
		}
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}

		ok,err,userId,appId:=util.VerifyAccessToken(token);

		if !ok ||err!=nil{
			app.Response(c,err,400)
			c.Abort()
			return
		}

		c.Set("userId",userId)
		c.Set("appId",appId)
		c.Next()
	}
}