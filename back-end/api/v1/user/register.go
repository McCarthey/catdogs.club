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
	passwdByte := []byte(user.Password)
	passwd := fmt.Sprintf("%x", md5.Sum(passwdByte))
	u := models.User{
		Email:    user.Email,
		Password: passwd,
	}
	u.Set()
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
