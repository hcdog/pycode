package dal

import (
	"github.com/hcdog/pycode/gomall/app/payment/biz/dal/mysql"
	// "github.com/hcdog/pycode/gomall/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
