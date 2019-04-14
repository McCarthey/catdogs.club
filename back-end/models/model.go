package models

import (
	"catdogs.club/back-end/configs/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func initTables() {
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	db.AutoMigrate(&User{})
}

func InitModel() {
	var err error
	db, err = gorm.Open("mysql", configs.GetDbAddr())
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 初始化表
	initTables()
}
