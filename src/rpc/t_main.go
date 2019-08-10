package rpc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const (
	HelloServiceName="HelloService"
)

type HelloService struct {

}
func (p *HelloService) Hello (request string,reply *string) error{
	fmt.Println("accept rpc call")
	*reply = "hello:"+request
	return nil
}

func (p *HelloService) SendEmail(request string,reply *string) error  {
	fmt.Println("accept send email")

	*reply = "Had send email to "+request
	return nil
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
	SendEmail(request string, reply *string) error
}
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func CreateRpc_t()  {
	defer func() {
		recover()
	}()
	var rpcPort = "1234"
	fmt.Println("start rpc listen:",rpcPort)

	err:=RegisterHelloService(new(HelloService))
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	//listener, err := net.Listen("tcp", ":1234")
	//if err != nil {
	//	log.Fatal("ListenTCP error:", err)
	//}
	//
	//for {
	//	conn, err := listener.Accept()
	//	fmt.Println("incoming rpc")
	//	if err != nil {
	//		log.Fatal("Accept error:", err)
	//	}
	//

	//	//go rpc.ServeConn(conn)// 提供普通的rpc服务
	//	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 提供jsonrpc服务
	//}

	// http 协议上提供jsonrpc服务
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {

		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser:request.Body,
			Writer:writer,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))

	})
	http.ListenAndServe(":1234",nil)
}
