package models

const (
	IsActivate = 1 << iota // 账号是否激活
)

type User struct {
	Id           int    `xorm:"pk autoincr"`
	Email        string `xorm:"varchar(36)"`
	Flags        int
	Name         string `xorm:"varchar(32)"`
	Openid       string `xorm:"varchar(64)"`
	Password     string `xorm:"varchar(32)"`
	PhoneNum     string `xorm:"varchar(18)"`
	RegisterTime int    `xorm:"created"`
}

func (u *User) Set() error {
	_, err := Db.Insert(u)
	return err
}

func (u *User) Get() (has bool, err error) {
	has, err = Db.Get(u)
	return
}

type VerifyCode struct {
	Id        int    `xorm:"pk autoincr"`
	Email     string `xorm:"varchar(36)"`
	PhoneNum  string `xorm:"varchar(15)"`
	Code      string `xorm:"varchar(16)"`
	Timestamp int    `xorm:"created"`
}

func (v *VerifyCode) Set() error {
	_, err := Db.Insert(v)
	return err
}

func (v *VerifyCode) Get() (has bool, err error) {
	has, err = Db.Get(v)
	return
}
