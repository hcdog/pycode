package dal

import (
	"github.com/hcdog/pycode/gomall/app/product/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
