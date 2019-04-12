package models

import (
	"catdogs.club/back-end/configs/common"
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitModel() {
	db, err := gorm.Open("mysql", configs.GetDbAddr())
	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
