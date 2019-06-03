package middleware

import (
	"catdogs-admin/libs"
	"catdogs-admin/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil || token == "" {
		logging.Error("Invalid cookie", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -1004,
		})
		c.Abort()
		return
	}
	claims, err := libs.ParseToken(token)
	if err != nil {
		logging.Error("Invalid token", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -1005,
		})
		c.Abort()
		return
	}
	exp := claims.ExpiresAt
	now := time.Now().Unix()
	if exp-now < 0 {
		libs.Resp(libs.R{
			C:    c,
			Code: -1006,
		})
		c.Abort()
		return
	}
	c.Set("openid", claims.Openid)
	c.Next()
}
