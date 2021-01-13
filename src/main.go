package main

import (
	"os"
	"fmt"
	
	"controller"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
		
	router.GET("/users", controller.GetUsers)
	router.GET("/users/{id:[0-9]+}", controller.GetUser)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/{id:[0-9]+}", controller.UpdateUser)
	router.DELETE("/users/{id:[0-9]+}", controller.DeleteUser)

	fasthttp.ListenAndServe(
		fmt.Sprintf(":%s", os.Getenv("PORT")),
		router.Handler,
	)
}