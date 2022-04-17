package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb",
		beego.AppConfig.String("mysql::User"),
		beego.AppConfig.String("mysql::Password"),
		beego.AppConfig.String("mysql::Host"),
		beego.AppConfig.String("mysql::Port"),
	)
	fmt.Println(DSN)
	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}
}
