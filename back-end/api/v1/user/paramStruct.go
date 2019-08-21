package user

type User struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}

type ProfileParam struct {
	Profile *Profile `form:"profile" json:"profile"`
}

type Profile struct {
	Name     string `form:"name" json:"name"`
	Gender   string `form:"gender" json:"gender"`
	Age      int32  `form:"age" json:"age"`
	PhoneNum string `form:"phonenum" json:"phonenum"`
	Email    string `form:"email" json:"email"`
	Birthday string `form:"birthday" json:"birthday"`
	City     string `form:"city" json:"city"`
	Address  string `form:"address" json:"address"`
}
