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
	err = saveUser(&user)
	if err != nil {
		fmt.Println(err)
	}
	code := libs.RandString(6)
	cont := fmt.Sprintf("您的邮箱验证码为: %s", code)
	go libs.SendMail(user.Email, "邮箱验证", cont)
	libs.Resp(c, 0, "success", gin.H{})
}

// 保存用户到数据库
func saveUser(param *User) error {
	pwData := md5.Sum([]byte(param.Password))
	pwS := fmt.Sprintf("%x", pwData)
	u := models.User{
		Email:        param.Email,
		Password:     pwS,
		RegisterTime: int(time.Now().Unix()),
	}
	err := u.Set()
	return err
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
