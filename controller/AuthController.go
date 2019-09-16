package controller

import (
	"gin-practice/filter"
)

var (
	AuthFilter *filter.AuthFilter
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Initialise() {
	AuthFilter = filter.NewAuthFilter(c.GinContext)
}

func (c *AuthController) Login() {
	userInfo, err := AuthFilter.Login()

	if err != nil {
		c.Response(FailedCode, err.Error(), nil)
		return
	}

	c.Response(SuccessCode, "success", userInfo)
}

func (c *AuthController) Register() {
	err := AuthFilter.Register()

	if err != nil {
		c.Response(FailedCode, err.Error(), nil)
		return
	}

	c.Response(SuccessCode, "success", nil)
}

func (c *AuthController) BulkInsertToES() {
	//req := elastic.NewBulkIndexRequest()
	//bulkResponse, err := esClient.EsConn.Bulk().Add().Do(context.Background())
}
