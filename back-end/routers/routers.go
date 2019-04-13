package routers

import (
	"catdogs.club/back-end/api/v1/test"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	registerApi(e)

	return e
}

func registerApi(e *gin.Engine) {
	apiv1 := e.Group("v1")
	apiv1.GET("/", test.Hello)
	apiv1.GET("/user", test.GetUser)
}
