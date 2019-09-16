package controller

import (
	"gin-practice/pkg/controllerHelper"
	"gin-practice/pkg/response"
	"net/http"
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

func (c *BaseController) SuccessResponse(data interface{}) {
	c.GinContext.JSON(http.StatusOK, &response.Response{RunTime: time.Since(c.Runtime).Seconds(), Code: 1, Message: "success", Data: data})
}

func (c *BaseController) FailResponse(data interface{}, err error) {
	c.GinContext.JSON(http.StatusOK, &response.Response{RunTime: time.Since(c.Runtime).Seconds(), Code: -1, Message: err.Error(), Data: data})
}
