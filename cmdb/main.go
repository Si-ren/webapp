package main

import (
	"github.com/astaxie/beego"
	_ "webapp/cmdb/models"
	_ "webapp/cmdb/routers"
)

func main() {
	beego.Run()
}
