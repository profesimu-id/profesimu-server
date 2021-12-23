package main

import (
	"profesimu/config"
	"profesimu/controllers"
	"profesimu/repository"
	"profesimu/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.InitDatabase()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	jwtService     services.JWTService        = services.NewJWTService()
	authService    services.AuthService       = services.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
)

func main() {
	config.InitDatabase()
	defer config.CloseDatabase(db)
	router := gin.Default()

	auth := router.Group("api/auth")
	// main := router.Group("api")

	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)

	router.Run()
}
