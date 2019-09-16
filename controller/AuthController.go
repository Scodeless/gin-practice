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
		c.FailResponse(nil, err)
		return
	}

	c.SuccessResponse(userInfo)
}

func (c *AuthController) Register() {
	err := AuthFilter.Register()

	if err != nil {
		c.FailResponse(nil, err)
		return
	}

	c.SuccessResponse(nil)
}

func (c *AuthController) BulkInsertToES() {
	//req := elastic.NewBulkIndexRequest()
	//bulkResponse, err := esClient.EsConn.Bulk().Add().Do(context.Background())
}
