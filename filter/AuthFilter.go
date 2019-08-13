package filter

import (
	"errors"
	"gin-practice/model/user"
	"gin-practice/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

var (
	AuthService *service.AuthService
)

type AuthFilter struct {
	Gin *gin.Context
}

func (f *AuthFilter) Login() (user []*user.User, err error) {
	userName := f.Gin.PostForm("username")
	password := f.Gin.PostForm("password")
	valid := validation.Validation{}
	valid.Required(userName, "username").Message("username is required")
	valid.Required(password, "password").Message("password is required")
	if valid.HasErrors() {
		return nil, errors.New(valid.Errors[0].String())
	}

	userInfo, err := AuthService.Login(userName, password)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userInfo, nil
}

func (f *AuthFilter) Register() (err error) {
	userName := f.Gin.PostForm("username")
	password := f.Gin.PostForm("password")
	valid := validation.Validation{}
	valid.Required(userName, "username").Message("username is required")
	valid.Required(password, "password").Message("password is required")
	if valid.HasErrors() {
		return errors.New(valid.Errors[0].String())
	}
	res, bools := AuthService.Register(userName, password)
	if !bools {
		return errors.New(res.Error())
	}
	return nil
}