package service

import (
	"gin-practice/model/user"
	"gin-practice/pkg/database"
)

type LoginService struct {

}

func (c *LoginService) Login(username, password string) (userInfo *user.User, err error) {
	userStruct := make([]*user.User, 0, 64)
	err = database.Conn.Table(user.GetTableName()).Select("user_id, age, gender, username").
		Where("username = ? and password = ?", username, password).Find(&userStruct).Error
	if err != nil {

	}

	return nil, nil
}