package main

import (
	_ "cmdb/models"
	_ "cmdb/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" //需要导入数据库所需 driver
)

func main() {
	beego.Run()
}
