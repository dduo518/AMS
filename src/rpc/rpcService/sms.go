package rpcService

import (
	"AMS/src/rpc/probuf"
	service "AMS/src/services"

)
import "golang.org/x/net/context"
import "errors"
type Sms struct{}

// 发送验证码
func (s *Sms) SendCode(ctx context.Context, req *probuf.SendCodeReq)(*probuf.SendCodeRes, error)  {
	if req.Phone==""{
		return nil,errors.New("Invalid Param")
	}
	result,err:=service.GetCodeWithLogin(req.Phone);
	if err!=nil{
		return nil,err
	}
	return &probuf.SendCodeRes{UserId: result.(string)}, nil
}

// 验证验证码
func (s *Sms) VerifyCode(ctx context.Context, req *probuf.VerifyCodeReq)(*probuf.VerifyCodeRes, error)  {
	if req.Phone=="" ||req.Code==""  {
		return nil,errors.New("Invalid Param")
	}
	result,err:=service.VerifyCode(req.Phone,req.Code,"appid")
	if err!=nil{
		return nil,err
	}
	return &probuf.VerifyCodeRes{
		AccessToken:result["AccessToken"].(string),
		UserId:result["UserId"].(string),
		Name:result["Name"].(string),
	},nil
}