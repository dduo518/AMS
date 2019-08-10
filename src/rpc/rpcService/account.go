package rpcService

import (
	"AMS/src/_interface"
	"AMS/src/rpc/probuf"
	service "AMS/src/services"
)

import "golang.org/x/net/context"
import "errors"
type Account struct{}

func (account *Account)LogUp(ctx context.Context, req *probuf.LogReq) (*probuf.LogRes, error) {
	var logupParams _interface.Account
	if req.Name=="" ||req.Email=="" ||req.Code==""||req.Phone==""||req.Password=="" {
		return nil,errors.New("Invalid Param")
	}
	logupParams.Phone=req.Phone
	logupParams.Name=req.Name
	logupParams.Code=req.Code
	logupParams.PassWord=req.Password
	logupParams.Type=req.Type
	logupParams.Address=req.Address
	logupParams.Email=req.Email
	logupParams.Type=req.Type
	result,err:= service.LogupServices(logupParams);

	if err!=nil{
		return nil,err
	}
	logupParams.UserId = result
	return &probuf.LogRes{
		UserId:result,
		Name:req.Name,
		Phone:req.Phone,
		Email:req.Email,
		Type:req.Type,
	},nil
}