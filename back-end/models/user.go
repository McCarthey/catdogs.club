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

func (u *User) Set() error {
	_, err := db.Insert(u)
	return err
}

func (u *User) Get() (has bool, err error) {
	has, err = db.Get(u)
	return
}

type VerifyCode struct {
	Id        int    `xorm:"pk autoincr"`
	Email     string `xorm:"varchar(36)"`
	Code      string `xorm:"varchar(8)"`
	Timestamp int
}

func (v *VerifyCode) Set() error {
	_, err := db.Insert(v)
	return err
}

func (v *VerifyCode) Get() (has bool, err error) {
	has, err = db.Get(v)
	return
}
