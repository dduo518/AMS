package model

import "github.com/jinzhu/gorm"

type AccountType struct {
	gorm.Model
	Name string `gorm:"INDEX"`
	Note string
	Type string `gorm:"NOT NULL"`
	AppId string `gorm:"NOT NULL"`
	AppKey string `gorm:"NOT NULL"`
}