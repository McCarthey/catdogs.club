package test

import (
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	u := &models.User{
		Name: "Yoko",
	}
	u.Get()
	c.JSON(http.StatusOK, gin.H{
		"name":  u.Name,
		"email": u.Email,
	})
}
