package interceptor

import (
	"AMS/src/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import "google.golang.org/grpc/codes"       // grpc 响应状态码
import "google.golang.org/grpc/metadata" // grpc metadata包
// auth 验证Token
func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid  string
		accessToken string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["token"]; ok {
		accessToken = val[0]
	}
	//const sqlSelect = "appKey"
	var account model.AccountType
	model.Engine.Model(
		&model.AccountType{
		}).Where(
		"app_id = ?", appid).Find(&account)
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(account.AppKey), nil
	})
	if err!=nil{
		return grpc.Errorf(codes.Unauthenticated, "Token认证信息失败")
	}
	if token.Valid{
		context.WithValue(ctx, "type", account.Type)
		context.WithValue(ctx, "accountTypeId", account.ID)
		return nil
	}

	return grpc.Errorf(codes.Unauthenticated, "Token认证信息失败")
}

// 注册interceptor
func AuthInterceptor(ctx context.Context,req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	err = auth(ctx)
	if err != nil {
		return
	}
	// 继续处理请求
	return handler(ctx, req)
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(resp interface{}, err error)  {
		log.Printf("gRPC method: %s,%v",info.FullMethod,req)
		resp,err =handler(ctx,req)
		fmt.Println(info.Server)
		log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
		return resp,err
}
