package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (u User) TableName() string {
	return "user"
}

// 这里的db为什么要上下文参数？(mysql.init后传过来的），可以在上层模块取消传递信号，好像有点牛笔
func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}
