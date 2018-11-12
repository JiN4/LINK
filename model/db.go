package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	// return config.Env("dialect", config.Env("datasource"))
	return "mysql", "root:password@tcp(kyotohack2018-f-1.cm68cjyvpekc.ap-northeast-1.rds.amazonaws.com:3306)/kyotohack2018_f?charset=utf8&parseTime=True&loc=Local"
}
