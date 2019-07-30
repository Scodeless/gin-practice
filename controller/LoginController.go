package controller

import (
	"gin-practice/pkg/response"
	"github.com/gin-gonic/gin"
)

var Response *response.Response

func Login(g *gin.Context) {
	userName := g.PostForm("username")
	password := g.PostForm("password")
	userInfo := make(map[string]string, 0)
	userInfo["username"] = userName
	userInfo["password"] = password
	Response.Response(g, 200, "success", "sajdflajsdlfja")
}