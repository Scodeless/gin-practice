package controller

import (
	"gin-practice/filter"
	"gin-practice/pkg/response"
	"github.com/gin-gonic/gin"
)

var (
	LoginFilter *filter.LoginFilter
)

type LoginController struct {
	Gin *gin.Context
	Res *response.Response
}

func (c *LoginController) Login() {
	_ = LoginFilter.Login()
	userName := c.Gin.PostForm("username")
	password := c.Gin.PostForm("password")
	userInfo := make(map[string]string, 0)
	userInfo["username"] = userName
	userInfo["password"] = password
	c.Res.Response(200, "success", "test")
}

func NewLoginController(g *gin.Context) LoginController {
	return LoginController{
		Gin: g,
		Res: &response.Response{G:g},
	}
}
