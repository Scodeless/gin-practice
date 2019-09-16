package controller

import (
	"gin-practice/pkg/controllerHelper"
	"gin-practice/pkg/response"
	"time"
)

const (
	SuccessCode = 1
	FailedCode  = -1
)

type BaseController struct {
	Runtime time.Time
	controllerHelper.Controller
}

type InitialiseInterface interface {
	Initialise()
}

func (c *BaseController) Prepare() {
	c.Runtime = time.Now()
	if app, ok := c.AppController.(InitialiseInterface); ok {
		app.Initialise()
	}
}

func (c *BaseController) Finish() {

}

func (c *BaseController) Response(code int, message string, data interface{}) {
	c.GinContext.JSON(200, &response.Response{RunTime: time.Since(c.Runtime).Seconds(), Code: code, Message: message, Data: data})
}
