package controller

import (
	"gin-practice/filter"
	"gin-practice/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginController struct {
	Gin *gin.Context
	Res *response.Response
	LoginFilter *filter.LoginFilter
}

func NewLoginController(g *gin.Context) *LoginController {
	return &LoginController{
		Gin: g,
		Res: &response.Response{G:g, Time:time.Now()},
		LoginFilter: &filter.LoginFilter{g},
	}
}

func (c *LoginController) Login() {

	userInfo, err := c.LoginFilter.Login()

	if err != nil {
		c.Res.Response(404, err.Error(), nil)
		return
	}

	c.Res.Response(200, "success", userInfo)
}

