package dal

import (
	"github.com/hcdog/pycode/gomall/app/checkout/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
