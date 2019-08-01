package filter

import (
	"errors"
	"gin-practice/model/user"
	"gin-practice/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

var (
	LoginService *service.LoginService
)

type LoginFilter struct {
	Gin *gin.Context
}

func (f *LoginFilter) Login() (user *user.User, err error) {
	userName := f.Gin.PostForm("username")
	password := f.Gin.PostForm("password")
	valid := validation.Validation{}
	valid.Required(userName, "username").Message("username is required")
	valid.Required(password, "password").Message("password is required")
	if valid.HasErrors() {
		return nil, errors.New(valid.Errors[0].String())
	}

	userInfo, err := LoginService.Login(userName, password)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userInfo, nil
}