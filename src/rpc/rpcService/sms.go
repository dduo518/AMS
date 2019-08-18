package rpcService

import (
	"AMS/src/rpc/probuf"
	service "AMS/src/services"
	"fmt"
	"log"
)
import "golang.org/x/net/context"
import "errors"
import "google.golang.org/grpc/metadata" // grpc metadata包

type Sms struct{}

func getMetadata(ctx context.Context) (string, string) {
	var (
		appid  string
		appkey string
	)
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md)
	if val,ok:=md["appid"];ok{
		log.Print(val)
		appid =val[0]
	}
	if val,ok:=md["appkey"];ok{
		log.Print(val)
		appkey =val[0]
	}
	return appid,appkey
}


// 发送验证码
func (s *Sms) SendCode(ctx context.Context, req *probuf.SendCodeReq)(*probuf.SendCodeRes, error)  {
	appid,_:=getMetadata(ctx)
	fmt.Println("send code ",appid)
	if req.Phone=="" || appid==""{
		return nil,errors.New("Invalid Param")
	}
	result,err:=service.GetCodeWithLogin(req.Phone,appid);
	if err!=nil{
		return nil,err
	}
	return &probuf.SendCodeRes{UserId: result.(string)}, nil
}

// 验证验证码
func (s *Sms) VerifyCode(ctx context.Context, req *probuf.VerifyCodeReq)(*probuf.VerifyCodeRes, error)  {
	appid,_:=getMetadata(ctx)
	if req.Phone=="" ||req.Code==""  {
		return nil,errors.New("Invalid Param")
	}
	result,err:=service.VerifyCode(req.Phone,req.Code,appid)
	if err!=nil{
		return nil,err
	}
	return &probuf.VerifyCodeRes{
		AccessToken:result["AccessToken"].(string),
		UserId:result["UserId"].(string),
		Name:result["Name"].(string),
	},nil
}