package user

type User struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}
