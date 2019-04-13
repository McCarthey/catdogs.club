package test

import (
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	u := new(models.User)
	u.Name = "Hejie"
	has, err := u.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"请求错误": "xxx"})
	}
	if has {
		c.JSON(http.StatusOK, gin.H{
			"UserName": u.Name,
			"Email":    u.Email,
		})
	}
}
