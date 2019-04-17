package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}

// @Summary 登录接口
// @Produce json
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {string} json "{"code": 0, "msg": "sucess", "data": {}}"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
