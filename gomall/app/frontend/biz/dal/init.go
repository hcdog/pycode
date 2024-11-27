package dal

import (
	"github.com/hcdog/pycode/gomall/app/frontend/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
