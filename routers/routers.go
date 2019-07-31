package routers

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.POST("/login", func(ctx *gin.Context) {
		LoginController := controller.NewLoginController(ctx)
		LoginController.Login()
	})

	return router
}