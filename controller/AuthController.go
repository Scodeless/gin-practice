package controller

import (
	"gin-practice/filter"
	"gin-practice/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthController struct {
	Gin *gin.Context
	Res *response.Response
	AuthFilter *filter.AuthFilter
	//BaseController
}

func NewAuthController(g *gin.Context) *AuthController {
	return &AuthController{
		Gin: g,
		Res: &response.Response{G:g, Time:time.Now()},
		AuthFilter: &filter.AuthFilter{Gin: g},
	}
}

func (c *AuthController) Login() {
	userInfo, err := c.AuthFilter.Login()

	if err != nil {
		c.Res.Response(FailedCode, err.Error(), nil)
		return
	}

	c.Res.Response(SuccessCode, "success", userInfo)
}

func (c *AuthController) Register() {
	err := c.AuthFilter.Register()

	if err != nil {
		c.Res.Response(FailedCode, err.Error(), nil)
	}

	c.Res.Response(SuccessCode, "success", nil)
}

