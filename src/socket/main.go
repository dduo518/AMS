package socket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)
var server * socketio.Server

func StartSocket(router *gin.Engine) {
	log.Print("StartSocket")
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Print(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	log.Println("Socket Serving at localhost:8000...")
	http.ListenAndServe(":8000", nil)
	//router.GET("/socket.io/", gin.WrapH(server))
	//router.GET("/socket.io/", func(c *gin.Context) {
	//	server.ServeHTTP (c.Writer,c.Request )
	//})
}
