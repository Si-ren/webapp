package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//
//var (
//	DB  *gorm.DB
//	err error
//)
//
//func init() {
//
//	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb?parseTime=true",
//		beego.AppConfig.String("mysql::User"),
//		beego.AppConfig.String("mysql::Password"),
//		beego.AppConfig.String("mysql::Host"),
//		beego.AppConfig.String("mysql::Port"),
//	)
//	fmt.Println(DSN)
//	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic(err.Error())
//	}
//
//}

var (
	mysql orm.Ormer
)

func init() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb?parseTime=true",
		beego.AppConfig.String("mysql::User"),
		beego.AppConfig.String("mysql::Password"),
		beego.AppConfig.String("mysql::Host"),
		beego.AppConfig.String("mysql::Port"),
	)
	fmt.Println(DSN)
	//注册驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL) //orm已经注册了，可省略
	fmt.Println("orm.RegisterDriver", err)
	//注册数据库,必须有一个叫default的别名
	err = orm.RegisterDataBase("default", "mysql", DSN, 30)
	fmt.Println("orm.RegisterDataBase", err)

	if db, err := orm.GetDB("default"); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	orm.RegisterModel(&User{})
	orm.RunSyncdb("default", false, true)

	mysql = orm.NewOrm()
}
