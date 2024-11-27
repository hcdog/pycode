package dal

import (
	"github.com/hcdog/pycode/gomall/app/user/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
