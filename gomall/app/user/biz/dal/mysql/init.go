package mysql

import (
	// "fmt"
	// "os"

	"github.com/hcdog/pycode/gomall/app/user/biz/model"
	// "github.com/hcdog/pycode/gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)



func Init() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,os.Getenv("MYSQL_USE"),os.Getenv("MYSQL_PASSWORD"),os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
}