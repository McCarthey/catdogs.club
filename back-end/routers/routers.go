package routers

import (
	"catdogs.club/back-end/api/v1/user"
	"catdogs.club/back-end/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	docs.SwaggerInfo.Title = "CatDogs API Doc"
	docs.SwaggerInfo.Description = "catdogs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "catdogs.club"
	docs.SwaggerInfo.BasePath = "api"
	e.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerApi(e)

	return e
}

func registerApi(e *gin.Engine) {
	apiv1 := e.Group("/api")
	apiv1.POST("/login", user.Login)
	apiv1.POST("/register", user.Register)
}
