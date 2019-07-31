package model

type User struct {
	UserId uint64 `gorm:"primary_key"`
	UserName string
	Age uint64
	Gender uint64
}

func (u User) TableName() string {
	return "users"
}
