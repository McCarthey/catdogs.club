package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestParam struct {
	Name string `form:"name"`
	Id   string `form:"id"`
}

func GetUser(c *gin.Context) {
	var testParam TestParam
	if err := c.ShouldBind(&testParam); err != nil {
		fmt.Println(err)
	}
	fmt.Println(testParam)
	c.JSON(http.StatusOK, testParam)
}
