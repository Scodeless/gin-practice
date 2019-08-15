package service

import (
	"fmt"
	"gin-practice/model/user"
	"gin-practice/pkg/db"
)

type AuthService struct {

}

/**
	登录方法
 */
func (c *AuthService) Login(username, password string) (userInfo []*user.User, err error) {
	userStruct := make([]*user.User, 0, 64)
	err = db.Conn.Table(user.GetTableName()).Select("user_id, age, gender, user_name").
		Where("user_name = ? and password = ?", username, password).Find(&userStruct).Error
	if err != nil {
		return nil, fmt.Errorf("query user failed, %v", err)
	}

	return userStruct, nil
}

func (c *AuthService) Register(username, password string) (err error) {
	userStructBuf := make([]*user.User, 0, 64)
	err = db.Conn.Table(user.GetTableName()).Select("user_id, age, gender, user_name").
		Where("user_name = ?", username).Find(&userStructBuf).Error
	if err != nil {
		return fmt.Errorf("query user_name exist error, %v", err)
	}
	if len(userStructBuf) > 0 {
		return fmt.Errorf("username is exist, please change another name")
	}
	//create new user
	userInfo := &user.User{UserName:username, Password:password}
	err = db.Conn.Table(user.GetTableName()).Create(&userInfo).Error
	if err != nil {
		return fmt.Errorf("insert to table error, %v", err)
	}
	return nil
}