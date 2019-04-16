package user

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"catdogs.club/back-end/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
	}
	pwData := md5.Sum([]byte(user.Password))
	pwS := fmt.Sprintf("%x", pwData)
	u := models.User{
		Email:    user.Email,
		Password: pwS,
	}
	u.Set()
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
