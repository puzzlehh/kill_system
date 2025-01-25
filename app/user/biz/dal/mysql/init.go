package mysql

import (
	"github.com/puzzlehh/kill_system/app/user/biz/model"
	"github.com/puzzlehh/kill_system/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	// 自动迁移
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
