package models

import (
	configs "catdogs.club/back-end/configs/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func InitModel() {
	var err error
	db, err = xorm.NewEngine("mysql", configs.DbAddr)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)

	initTables()
}

func initTables() {
	db.Sync2(new(User))
}
