package main

import (
	"gin-test/controller"
	"gin-test/pkg/database"
	"gin-test/pkg/rpc/server"
	"github.com/gin-gonic/gin"
)

func init()  {
	//协程开启rpc服务
	go server.StartRpc()

	//连接数据库
	database.DbInit()
}

func main()  {

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login",  controller.Login)

	_ = r.Run(":3000")
}