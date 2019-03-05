package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/goalong/bingo/conf"

)

var (
	// gorm mysql db connection
	db *gorm.DB
)



func init() {
	var err error
	// mysql conn
	config := conf.Config

	db, err = gorm.Open("mysql", config.DB.User+":"+config.DB.Password+
		"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name+
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=90s")
	if err != nil {
		log.Fatal("connect db fail")
	}
	log.Println("connect db succcess")
	db.SingularTable(true) // 全局禁用表名复数


	// gorm debug log
	if config.APP.Debug {
		db.LogMode(true)
	}
}
