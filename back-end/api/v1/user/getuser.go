package user

import (
	"catdogs.club/back-end/libs"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	libs.Resp(c, 0, "success", gin.H{})
}
