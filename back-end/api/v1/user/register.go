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
	pwData := md5.Sum([]byte(user.Password))
	pwS := fmt.Sprintf("%x", pwData)
	u := models.User{
		Email:        user.Email,
		Password:     pwS,
		RegisterTime: int(time.Now().Unix()),
	}
	u.Set()
	libs.Resp(c, 0, "success", gin.H{})
}

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
