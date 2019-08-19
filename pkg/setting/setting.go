package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Ini *ini.File
	RunMode string
	Port int
	Db struct{
		   Host string
		   Port int
		   User string
		   Name string
		   TablePrifix string
		   Password string
	   }
	)

func init() {
	var err error
	Ini, err = ini.Load("conf/app.conf")
	if err !=nil {
		log.Fatalf("fail to load config file: %v", err)
	}
	LoadBase()
	LoadServer()
	LoadDataBase()
}

func LoadBase() {
	RunMode = Ini.Section("").Key("RUN_MODE").MustString("dev")
}

func LoadServer() {
	server, err := Ini.GetSection("server")

	if err != nil {
		log.Fatalf("fail to load section server,please add server config")
	}

	Port = server.Key("HTTP_PORT").MustInt(8080)
}

func LoadDataBase() {
	database, err := Ini.GetSection("database")

	if err != nil {
		log.Fatalf("fail to load section db,please add db config")
	}
	Db.Host = database.Key("HOST").MustString("127.0.0.1")
	Db.Port = database.Key("PORT").MustInt(3306)
	Db.User = database.Key("USER").MustString("root")
	Db.Password = database.Key("PASSWORD").MustString("")
	Db.Name = database.Key("NAME").MustString("")
	Db.TablePrifix = database.Key("TABLE_PRIFIX").MustString("")
	//Db{
	//	Host: db.Key("HOST").MustString("127.0.0.1"),
	//	Port: db.Key("PORT").MustInt(3306),
	//	User: db.Key("USER").MustString("root"),
	//	Password: db.Key("PASSWORD").MustString(""),
	//	Name: db.Key("NAME").MustString(""),
	//	TablePrifix: db.Key("TABLE_PRIFIX").MustString(""),
	//}
}