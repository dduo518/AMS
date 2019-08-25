package app

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context,result interface{},code int)   {
	if code==0{
		c.JSON(GetHttpStatus(code),gin.H{
			"status":GetStatus(code),
			"messgae":GetMsg(code),
			"code":code,
			"data":result,
		})
	}else{

		c.JSON(GetHttpStatus(code),gin.H{
			"status":GetStatus(code),
			"messgae":result,
			"code":code,
			"data":result,
		})
	}
	c.Abort()
	return
}