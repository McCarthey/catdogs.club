package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email        string
	Flags        int
	Name         string
	Openid       string
	Password     string
	PhoneNum     string
	RegisterTime int
}

func (u *User) Set() {
	db.Create(&u)
}

func (u *User) Get() {
	db.Find(&u)
}
