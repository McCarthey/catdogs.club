package user

import (
	"catdogs.club/back-end/logging"
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
		logging.Error(err)
	}
	logging.Info("v param", vparam)
	v := models.VerifyCode{
		Code: vparam.V,
	}
	has, err := v.Get()
	logging.Info("verify code: ", v)
	if has {
		sql := "update user set flags=flags|? where email=?;"
		_, err := models.Db.Exec(sql, models.IsActivate, v.Email)
		if err != nil {
			c.String(0, err.Error(), "")
			return
		}
	}
	c.String(0, "账号激活成功", "")
}
