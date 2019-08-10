package rpc

import (
	"AMS/config"
	"AMS/src/rpc/interceptor"
	"AMS/src/rpc/probuf"
	"AMS/src/rpc/rpcService"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)



func CreateRpc()  {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor.AuthInterceptor))
	newServer := grpc.NewServer(opts...) //创建gRPC服务
	//
	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	probuf.RegisterSmsServer(newServer, &rpcService.Sms{})
	probuf.RegisterAccountServer(newServer, &rpcService.Account{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(newServer)
	var port = config.Conf.GprcConf.Port
	var network = config.Conf.GprcConf.Network
	fmt.Println("rpc 监听:", port)
	// 将监听交给gRPC服务处理
	lis, err := net.Listen(network, port )  //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
		return
	}

	err = newServer.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}