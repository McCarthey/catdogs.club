package main

import (
	"catdogs.club/back-end/models"
	"catdogs.club/back-end/routers"
	"fmt"
	"net/http"
	"time"
)

func init() {
	models.InitModel()
}

func main() {
	u := &models.User{
		Name:     "Yoko",
		Email:    "18836617@qq.com",
		PhoneNum: "15336200123",
	}
	u.Set()
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

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
