package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController interface repre middleware
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

//NewAuthController implements auth
func NewAuthController() AuthController {
	return &authController{}
}
func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello login"})
}
func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello register"})
}
