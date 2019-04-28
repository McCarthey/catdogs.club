package user

import (
	"catdogs.club/back-end/logger"
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
)

type VParam struct {
	V string `form:"v"`
}

func Verify(c *gin.Context) {
	var vparam VParam
	err := c.ShouldBind(vparam)
	if err != nil {
		logger.Error(err)
	}
	v := models.VerifyCode{
		Code: vparam.V,
	}
	has, err := v.Get()
	if has {
		updateUserFlag(&v)
	}
	c.String(0, "账号激活成功", "")
}

func updateUserFlag(v *models.VerifyCode) {
	u := models.User{
		Email: v.Email,
	}
	has, err := u.Get()
	if err != nil {
		logger.Error(err)
	}
	if has {
		u.Flags |= models.IsActivate
		models.UpdateByEmail(&u)
	}
}
