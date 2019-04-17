package user

import (
	"crypto/md5"
	"fmt"
	"time"

	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
	}

	has := verifyUser(&user)
	if has {
		libs.Resp(c, -1000, "用户已存在", gin.H{})
		return
	}
	code := libs.RandString(6)
	cont := fmt.Sprintf("您的邮箱验证码为: %s", code)
	go saveUser(&user)
	go saveCode(code, &user)
	go libs.SendMail(user.Email, "邮箱验证", cont)
	libs.Resp(c, 0, "success", gin.H{})
}

func saveCode(code string, param *User) {
	v := models.VerifyCode{
		Email:     param.Email,
		Code:      code,
		Timestamp: int(time.Now().Unix()),
	}
	err := v.Set()
	if err != nil {
		fmt.Println(err)
	}
}

// 保存用户到数据库
func saveUser(param *User) {
	pwData := md5.Sum([]byte(param.Password))
	pwS := fmt.Sprintf("%x", pwData)
	u := models.User{
		Email:        param.Email,
		Password:     pwS,
		RegisterTime: int(time.Now().Unix()),
	}
	err := u.Set()
	if err != nil {
		fmt.Println(err)
	}
}

// 验证用户是否存在
func verifyUser(param *User) bool {
	u := models.User{
		Email: param.Email,
	}
	has, err := u.Get()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return has
}
