package main

import (
	"profesimu/config"
	"profesimu/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.InitDatabase()
	authController controllers.AuthController = controllers.NewAuthController()
)

func main() {
	defer config.CloseDatabase(db)
	router := gin.Default()

	auth := router.Group("api/auth")
	// main := router.Group("api")

	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)

	router.Run()
}
