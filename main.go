package main

import (
	"fmt"
	"gin-practice/controller"
	"gin-practice/pkg/rpc/server"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

func init()  {
	//协程开启rpc服务
	go server.StartRpc()

	//连接数据库
	//database.DbInit()

	fmt.Println(beego.BConfig.AppName)
}

func main()  {

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login",  controller.Login)

	_ = r.Run(":3000")
}