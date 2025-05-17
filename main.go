package main

import (
	"github.com/gin-gonic/gin"
	"wallet/the_wallet/handler"
)

func main() {
	router := gin.Default()

	router.GET("/hello", handler.Handle)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
