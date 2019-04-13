package models

type Profile struct {
	NickName string
	Country  string
}

type User struct {
	Email        string `xorm:"varchar(50)"`
	Flags        int
	Name         string
	Openid       string
	Password     string
	PhoneNum     string
	RegisterTime int
}

func (u *User) Set() error {
	_, err := engine.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get() (bool, error) {
	has, err := engine.Get(u)
	return has, err
}
