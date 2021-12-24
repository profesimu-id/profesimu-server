package main

import (
	"profesimu/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()
	router := gin.Default()

	auth := router.Group("api/auth")
	// main := router.Group("api")

	auth.POST("/login")
	auth.POST("/register")

	router.Run()
}
