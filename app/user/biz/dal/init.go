package dal

import (
	"kill_system/app/user/biz/dal/mysql"
	"kill_system/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
