package model

import (
	"AMS/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
import "github.com/jinzhu/gorm"
var Engine *gorm.DB

func init()  {
	var url = config.Conf.Mysql.Username + ":" +
		config.Conf.Mysql.Password + "@tcp("+config.Conf.Mysql.Host+":"+config.Conf.Mysql.Port+")/" +
		config.Conf.Mysql.Database
	url = url + `?charset=utf8&parseTime=True&loc=Local`
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if config.Conf.GO_ENV == "dev" {
		Engine = db.Debug()
	} else {
		Engine = db.Debug()
	}
	if !db.HasTable(&Account{}) {
		db.CreateTable(&Account{})
	}

	if !db.HasTable(&AccountType{}) {
		db.CreateTable(&AccountType{})
	}
}

func Close()  {
	defer Engine.Close()
}