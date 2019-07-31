package login

import (
	"gin-practice/filter/login"
	"gin-practice/pkg/response"
	"github.com/gin-gonic/gin"
)

var (

	Filter *login.Filter
)

func Login(c *gin.Context) {
	_ = Filter.Login()
	Res := response.Response{ c}
	userName := c.PostForm("username")
	password := c.PostForm("password")
	userInfo := make(map[string]string, 0)
	userInfo["username"] = userName
	userInfo["password"] = password
	Res.Response(200, "success", "sajdflajsdlfja")
}