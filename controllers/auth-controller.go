package controllers

import (
	"net/http"
	"profesimu/dto"
	"profesimu/entity"
	"profesimu/helper"
	"profesimu/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// masukan service yang dibutuhkan
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO

	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if value, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(value.ID), 10))
		value.Token = generateToken
		response := helper.BuildResponse(true, "OK", value)
		ctx.JSON(http.StatusOK, response)
	}

	errRes := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, errRes)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerData dto.RegisterDTO

	errDTO := ctx.ShouldBindJSON(&registerData)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if !c.authService.IsDuplicateEmail(registerData.Email) {
		errRes := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, errRes)
	} else {
		createdUser := c.authService.CreateUser(dto.UserCreateDTO(registerData))
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusOK, response)
	}

}
