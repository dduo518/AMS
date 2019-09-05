package socket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)
var server * socketio.Server

func StartSocket(router *gin.Engine) {
	log.Print("StartSocket")
	pt := polling.Default
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

	if err != nil {
		log.Print(err)
	}

	server.OnConnect("/", func(socket socketio.Conn) error {
		socket.SetContext("")
		fmt.Println("connected:", socket.ID())
		return nil
	})

	server.OnDisconnect("/", func(socket socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	go server.Serve()

	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/",  gin.WrapH(server))
}
