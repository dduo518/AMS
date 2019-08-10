package rpcService

import (
	"AMS/src/rpc/probuf"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
	"golang.org/x/net/context"
)

// 业务实现方法的容器
type Server struct{}

// 为server定义 DoMD5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *test.Req[相应接口定义的请求参数])
// 返回 (*test.Res[相应接口定义的返回参数，必须用指针], error)
func (s *Server) DoMD5(ctx context.Context, in *probuf.Req) (*probuf.Res, error) {
	fmt.Println("MD5方法请求JSON:"+in.JsonStr)
	fmt.Println("MD5方法请求Phone:"+in.Phone)
	var sed int
	sed, err := strconv.Atoi(in.JsonStr)
	if err!=nil{
		fmt.Println(err)
	}
	time.Sleep(time.Duration(sed)*time.Second)
	return &probuf.Res{BackJson: "MD5 :" + in.JsonStr}, nil
}

func (s *Server) GetAccessToken(ctx context.Context, in *probuf.Req) (*probuf.Res, error) {
	fmt.Println("GetAccessToken方法请求JSON:"+in.JsonStr)
	fmt.Println("GetAccessToken方法请求Phone:"+in.Phone)
	return &probuf.Res{BackJson: "accessToken :" + fmt.Sprintf("%x", md5.Sum([]byte(in.JsonStr)))}, nil
}

