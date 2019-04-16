package models

type User struct {
	Id           int    `xorm:"pk autoincr"`
	Email        string `xorm:"varchar(36)"`
	Flags        int
	Name         string `xorm:"varchar(32)"`
	Openid       string `xorm:"varchar(64)"`
	Password     string `xorm:"varchar(32)"`
	PhoneNum     string `xorm:"varchar(18)"`
	RegisterTime int
}

func (u *User) Set() {
	db.Insert(u)
}

func (u *User) Get() (has bool, err error) {
	has, err = db.Get(u)
	return
}
