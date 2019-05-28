package user

import (
	"catdogs.club/back-end/libs"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "", "118.24.146.34", true, true)
	libs.Resp(libs.R{
		C:    c,
		Code: 0,
	})
}
