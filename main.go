package main

import (
	"profesimu/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()
	router := gin.Default()

	auth := router.Group("/auth")
	main := router.Group("/")

	auth.POST("/login")
	auth.POST("/register")
	main.GET("/profession")

	router.Run()
}
