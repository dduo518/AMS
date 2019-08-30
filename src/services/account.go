package services

import (
	"AMS/src/_interface"
	"AMS/src/models"
	"AMS/src/utils"
	"errors"
	"strconv"
	"time"
	"AMS/config"
)
import _ "AMS/src/_interface"

func LogupServices(account _interface.Account) (string,error) {
	account.LogupInitAccount()
	var accountModle model.Account
	accountModle.PassWord = account.PassWord
	accountModle.UserId = account.UserId
	accountModle.Name = account.Name
	accountModle.Phone = account.Phone
	accountModle.Email = account.Email
	accountModle.Verify = account.Verify
	accountModle.AppId = account.AppId
	model.Engine.Create(&accountModle)
	return accountModle.UserId,nil
}

func CheckPhone(phone string) map[string]interface{} {
	var count int
	var account model.Account
	model.Engine.Model(&model.Account{}).Where("phone = ?", phone).Find(&account).Count(&count)
	result := map[string]interface{}{
		"islogup":true,
	}
	if count==0{
		result["islogup"] = false
		return result
	}
	result["user_id"] = account.UserId
	return result
}

func VerifyCode(phone,code,appid string)(map[string]interface{},error)  {
	var account _interface.Account
	const sqlSelect = "name,id,phone,code,user_id,code_exp"
	model.Engine.Model(
		&model.Account{
		}).Select(sqlSelect).Where(
			"phone = ?", phone).Find(&account)
	if account.UserId=="" {
		return nil, errors.New("用户不存在")
	}
	if config.Conf.GO_ENV=="prod"{
		if account.Code!=code{
			return nil,errors.New("验证码无效")
		}
		nowStamp :=time.Now().Unix()
		codeExp, _ := strconv.ParseInt(account.CodeExp, 10, 64)
		if nowStamp>codeExp{
			return nil,errors.New("验证码过期")
		}
	}

	model.Engine.Model(&model.Account{}).Where(
		"phone = ?",phone).Update(map[string]interface{}{
		"verify":true,
	})
	accessToken:=util.CreateToken(account.UserId,appid)
	var resp = map[string]interface{}{
		"UserId":account.UserId,
		"Name":account.Name,
		"AccessToken":accessToken,
	}
	return resp,nil
}