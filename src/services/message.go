package services

import (
	"AMS/src/_interface"
	"AMS/src/models"
	"math/rand"
	"strconv"
	"time"
)

// 发送验证码
func SendVerifierMessage(phone,code string) (interface{},error)  {
	var message _interface.Message
	result,err:=message.VerifierMessage(phone,code)
	return result,err
}

// 获取登录验证码，如果手机号是新的则创建用户
func GetCodeWithLogin(phone string)(interface{},error) {
	result := CheckPhone(phone);
	if val,ok:=result["islogup"];ok && val==true{
		userId,_:=result["user_id"]
		UpdateCode(phone)
		return userId,nil
	}
	var logupParams _interface.Account
	logupParams.Phone =phone
	logupParams.Name =phone
	logupParams.Verify = false
	userId,err:= LogupServices(logupParams);
	UpdateCode(phone)
	return userId,err
}

// 更新验证码
func UpdateCode(phone string) (interface{},error) {
	code:=rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	expTimeStamp:= time.Now().Unix()+60*3
	var account _interface.Account
	account.Phone = phone
	model.Engine.Model(&model.Account{}).Where("phone=?",phone).Update(map[string]interface{}{
		"code":strconv.Itoa(int(code)),
		"code_exp":expTimeStamp,
	})
	return SendVerifierMessage(phone,strconv.Itoa(int(code)))
}