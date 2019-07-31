package main

import (
	"gin-practice/pkg/rpc/server"
	"gin-practice/routers"
)

func init()  {
	//协程开启rpc服务
	go server.StartRpc()

	//连接数据库
	//database.DbInit()
}

func main()  {
	r := routers.RouterInit()
	_ = r.Run(":3000")
}