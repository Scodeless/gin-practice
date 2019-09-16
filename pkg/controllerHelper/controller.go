package controllerHelper

import "github.com/gin-gonic/gin"

type Controller struct {
	AppController interface{}
	GinContext    *gin.Context
}

func (c *Controller) Init(ctl ControllerInterface, ctx *gin.Context) {
	c.AppController = ctl
	c.GinContext = ctx
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}
