package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"../conf"

)

var (
	// gorm mysql db connection
	db *gorm.DB
)



func initDB() {
	var err error
	// mysql conn
	conf := conf.Config
	for {
		db, err = gorm.Open("mysql", conf.DB.User+":"+conf.DB.Password+
			"@tcp("+conf.DB.Host+":"+conf.DB.Port+")/"+conf.DB.Name+
			"?charset=utf8mb4&parseTime=True&loc=Local&timeout=90s")
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}

	// gorm debug log
	if conf.APP.Debug {
		db.LogMode(true)
	}
}
