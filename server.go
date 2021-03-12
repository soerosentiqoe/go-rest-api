package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soerosentiqoe/go-rest-api/config"
	"github.com/soerosentiqoe/go-rest-api/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDb()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run(":8081")

}
