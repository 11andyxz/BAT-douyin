package dal

import (
	"BAT-douyin/commen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlInit() {
	var err error
	dsn := commen.DSN
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
