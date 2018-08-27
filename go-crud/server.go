package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/muve-go/go-crud/auth"
	"github.com/lakshanwd/muve-go/go-crud/handler"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	auth.InitAuth()
	fmt.Println("handeling routes")
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.POST("/authenticate", handler.Authenticate)

	authorized := router.Group("/")
	authorized.Use(auth.JWT())
	{
		authorized.GET("/booking/:id", handler.BookingGet)
		authorized.GET("/booking", handler.BookingGet)
		authorized.POST("/booking", handler.BookingPost)
		authorized.PUT("/booking/:id", handler.BookingPut)
		authorized.DELETE("/booking/:id", handler.BookingDelete)
	}
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Content-Type", "Authorization")
	router.Use(cors.New(config))
	router.Run(":8080")
}
