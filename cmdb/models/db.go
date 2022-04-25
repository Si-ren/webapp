package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func init() {

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb?parseTime=true",
		beego.AppConfig.String("mysql::User"),
		beego.AppConfig.String("mysql::Password"),
		beego.AppConfig.String("mysql::Host"),
		beego.AppConfig.String("mysql::Port"),
	)
	fmt.Println(DSN)
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

}
