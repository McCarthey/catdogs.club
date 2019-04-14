package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) == nil {
		fmt.Println(user.Email)
		fmt.Println(user.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
