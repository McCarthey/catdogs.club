package test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloParam struct {
	Name string `form:"name"`
}

func Hello(c *gin.Context) {
	var p HelloParam
	if err := c.BindJSON(&p); err != nil {
		fmt.Println(err)
	}
	log.Println("============")
	c.JSON(http.StatusOK, gin.H{
		"Hello": p.Name,
	})
}
