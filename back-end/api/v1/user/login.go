package user

import (
	"crypto/md5"
	"fmt"

	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err.Error())
	}
	u := models.User{
		Email: user.Email,
	}
	has, err := u.Get()
	fmt.Println(err)
	if !has {
		libs.Resp(c, -1002, "用户不存在", gin.H{})
		return
	}
	pwd := md5.Sum([]byte(user.Password))
	pwdHex := fmt.Sprintf("%x", pwd)
	if pwdHex != u.Password {
		libs.Resp(c, -1003, "密码错误", gin.H{})
		return
	}
	libs.Resp(c, 0, "登录成功", gin.H{})
}
