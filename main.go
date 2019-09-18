package main

import (
	"fmt"
	"gin-practice/pkg/db"
	"gin-practice/pkg/rabbitMQ"
	"gin-practice/pkg/rpc/server"
	"gin-practice/pkg/setting"
	"gin-practice/routers"
	"log"
	"time"
)

func init() {
	//协程开启rpc服务
	go server.StartRpc()
}

func main() {
	fmt.Printf("app run in %s \n", setting.RunMode)

	//连接mysql数据库
	conn, err := db.GetDbConn()
	if err != nil {
		log.Fatalf("db connect failed, %v", err)
	}
	fmt.Println("mysql database connect success")
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	conn.SingularTable(true)
	conn.DB().SetMaxIdleConns(10)  //SetMaxIdleConns用于设置闲置的连接数
	conn.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	conn.DB().SetConnMaxLifetime(2 * time.Minute)
	defer conn.Close()

	//连接es
	//esConn, err := esClient.ConnectElasticSearchClient()
	//if err != nil {
	//	log.Fatalf("es connect failed, %v", err)
	//}
	//_, _, _ = esConn.Ping("http://127.0.0.1:9200").Do(context.Background())
	//fmt.Println("es database connect success")

	//连接RabbitMQ
	rmqConn := rabbitMQ.ConnectRabbitMQServer()
	defer rmqConn.Close()
	fmt.Println("rabbitMQ connect success")

	//初始化路由
	r := routers.RouterInit()
	//use pprof listen
	//ginpprof.Wrapper(r)
	_ = r.Run(":3000")
}
