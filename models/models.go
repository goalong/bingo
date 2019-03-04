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
	conf := conf.Config

	db, err = gorm.Open("mysql", conf.DB.User+":"+conf.DB.Password+
		"@tcp("+conf.DB.Host+":"+conf.DB.Port+")/"+conf.DB.Name+
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=90s")
	if err != nil {
		log.Fatal("connect db fail")
	}
	log.Println("connect db succcess")
	db.SingularTable(true) // 全局禁用表名复数


	// gorm debug log
	if conf.APP.Debug {
		db.LogMode(true)
	}
}
