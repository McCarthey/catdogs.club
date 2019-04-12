package models

type User struct {
	Email    string `gorm:"column:email"`
	Flags    int64  `gorm:"column:flags"`
	Name     string `gorm:"column:name"`
	Openid   string `gorm:"column:openid"`
	Password string `gorm:"column:password"`
	PhoneNum string `gorm:"column:phoneNum"`
}

func (u *User) SetUser() {
	db.Create(&u)
}
