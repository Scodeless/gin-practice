package routers

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {

	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.POST("/auth/login", func(ctx *gin.Context) {
			AuthController := controller.NewAuthController(ctx)
			AuthController.Login()
		})
		v1.POST("/auth/register", func(ctx *gin.Context) {
			AuthController := controller.NewAuthController(ctx)
			AuthController.Register()
		})
	}

	return router
}