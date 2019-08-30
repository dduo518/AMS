package main

import (
	"AMS/config"
	"AMS/src/models"
	"AMS/src/router"
	"AMS/src/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
)
func CreateWebServer(port string)  {
	defer func() {
		recover()
	}()
	server := gin.Default()
	defer model.Close()
	gin.SetMode(gin.ReleaseMode)
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	router.SetUp(server)
	fmt.Println("start web server listen:",port)
	server.Run(port)
}

func main()  {

	defer func() {
		recover()
	}()
	go CreateWebServer(config.Conf.PORT)

	go rpc.CreateRpc()


	select {
	}

}
