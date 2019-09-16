package controllerHelper

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	Init(ctl ControllerInterface, ctx *gin.Context)
	Prepare()
	Finish()
}
