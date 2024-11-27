package dal

import (
	"github.com/hcdog/pycode/gomall/app/cart/biz/dal/mysql"
	// "github.com/hcdog/pycode/gomall/app/cart/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
