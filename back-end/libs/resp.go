package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(c *gin.Context, code int, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
