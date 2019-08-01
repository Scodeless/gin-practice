package main

import (
	"gin-practice/pkg/database"
	"gin-practice/pkg/rpc/server"
	"gin-practice/routers"
	"log"
	"time"
)

func init()  {
	//协程开启rpc服务
	go server.StartRpc()
}

func main()  {
	//连接数据库
	db, err := database.DbInit()
	if err != nil {
		log.Fatalf("database connect failed, %v", err)
	}
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10) //SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetConnMaxLifetime(2 * time.Minute)
	defer db.Close()
	r := routers.RouterInit()
	//use pprof listen
	//ginpprof.Wrapper(r)
	_ = r.Run(":3000")
}