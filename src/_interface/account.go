package _interface

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/google/uuid"
)

type Account struct {
	Name string `json:"name"`
	UserId string `json:"user_id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role int `json:"role"`
	Address string `json:"address"`
	PassWord string `json:"password"`
	Type string `json:"type"`
	Verify bool `json:"verify"`
	Code string `json:"code"`
	CodeExp string `json:"code_exp"`
}


func (account *Account) LogupInitAccount()  {
	account.SetUuid()
	account.SetSha1()
}

func (account *Account) SetUuid(){
	var userId,_ = uuid.NewUUID()
	account.UserId = userId.String()
}

func (account *Account) SetSha1(){
	if account.PassWord==""{
		account.PassWord = "123456"
	}
	pwHash := sha1.New()
	pwHash.Write([]byte(account.PassWord))
	account.PassWord =  hex.EncodeToString(pwHash.Sum([]byte("")))

}