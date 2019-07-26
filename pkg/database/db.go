package database

import (
	"fmt"
	"gin-test/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

func DbInit()  {
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",setting.Db.User, setting.Db.Password, setting.Db.Host, setting.Db.Name)
	fmt.Println(conStr)
	db, err := gorm.Open("mysql", conStr)

	if err != nil {
		log.Fatalf("database connect fail, please checkout you config:%s", err)
	}

	defer db.Close()

	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)

	//设置默认表名的命名规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return setting.Db.TablePrifix + defaultTableName
	}

	db.DB().SetMaxIdleConns(10) //SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetConnMaxLifetime(2 * time.Minute)
}