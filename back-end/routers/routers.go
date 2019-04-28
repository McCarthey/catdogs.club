package routers

import (
	"catdogs.club/back-end/api/v1/user"
	configs "catdogs.club/back-end/configs/common"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	if configs.EnvModel == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	registerApi(e)

	return e
}

func registerApi(e *gin.Engine) {
	apiv1 := e.Group("/api")
	apiv1.GET("/verify", user.Verify)
	apiv1.POST("/login", user.Login)
	apiv1.POST("/register", user.Register)
}
