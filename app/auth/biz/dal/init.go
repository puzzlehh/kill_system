package dal

import (
	"github.com/puzzlehh/kill_system/app/auth/biz/dal/mysql"
	"github.com/puzzlehh/kill_system/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
