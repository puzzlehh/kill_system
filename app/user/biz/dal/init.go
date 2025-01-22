package dal

import (
	"github.com/puzzlehh/kill_system/app/user/biz/dal/mysql"
	"github.com/puzzlehh/kill_system/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
