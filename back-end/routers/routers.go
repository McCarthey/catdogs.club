package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	return e
}
