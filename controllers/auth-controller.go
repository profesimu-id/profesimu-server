package controllers

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// masukan service yang dibutuhkan

}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello register",
	})
}
