package user

import (
	"net/http"

	"catdogs.club/back-end/logging"
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	v := c.Query("v")
	vm := models.VerifyCode{
		Code: v,
	}
	has, err := vm.Get()
	if err != nil {
		logging.Error("vm get error: ", err)
	}
	logging.Info("verify code: ", v)
	if has {
		sql := "update user set flags=flags|? where email=?;"
		_, err := models.Db.Exec(sql, models.IsActivate, vm.Email)
		if err != nil {
			c.String(0, err.Error())
			return
		}
	}
	c.String(http.StatusOK, "账号激活成功\n")
}
