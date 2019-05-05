package user

import (
	"crypto/md5"
	"fmt"

	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/logging"
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
		logging.Info("bind user err", err.Error())
	}
	u := models.User{
		Email: user.Email,
	}
	logging.Info(u)
	has, err := u.Get()
	fmt.Println(err)
	if err != nil {
		logging.Info("get user err", err.Error())
		libs.Resp(libs.R{
			C:    c,
			Code: -999,
		})
		return
	}
	if !has {
		libs.Resp(libs.R{
			C:    c,
			Code: -1002,
		})
		return
	}
	pwd := md5.Sum([]byte(user.Password))
	pwdHex := fmt.Sprintf("%x", pwd)
	if pwdHex != u.Password {
		libs.Resp(libs.R{
			C:    c,
			Code: -1003,
		})
		return
	}
	libs.Resp(libs.R{
		C:    c,
		Code: 0,
		Msg:  "登录成功",
	})
}
