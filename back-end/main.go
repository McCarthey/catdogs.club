// @title CatDogs API
// @version 1.0
// @description CatDogs API 文档
// @contact.name Yoko
// @BasePath /api

package main

import (
	"fmt"
	"net/http"
	"time"

	_ "catdogs.club/back-end/docs"
	"catdogs.club/back-end/models"
	"catdogs.club/back-end/routers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	models.InitModel()
}

func main() {
	runServer()
}

func runServer() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    18 * time.Second,
		WriteTimeout:   18 * time.Second,
		MaxHeaderBytes: 1 << 28,
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
