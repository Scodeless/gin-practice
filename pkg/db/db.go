package db

import (
	"fmt"
	"gin-practice/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

func GetDbConn() (*gorm.DB, error) {
	//连接数据库
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",setting.Db.User, setting.Db.Password, setting.Db.Host, setting.Db.Name)
	//fmt.Println(conStr)
	db, err := gorm.Open("mysql", conStr)
	Conn = db

	//设置默认表名的命名规则
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return setting.Db.TablePrifix + defaultTableName
	//}

	return db, err
}