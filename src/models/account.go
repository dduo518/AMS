package model

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Name string `gorm:"INDEX"`
	UserId string `gorm:"UNIQUE;NOT NULL;PRIMARY_KEY";`
	Email string
	Phone string
	Role int `gorm:"NOT NULL"`
	Address string
	PassWord string `gorm:"NOT NULL"`
	Type string `gorm:"NOT NULL"`
	Verify bool `gorm:"DEFAULE:false"`
	Code string `json:"code"`
	CodeExp int32 `json:"code_exp"`
}
