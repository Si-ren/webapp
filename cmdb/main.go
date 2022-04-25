package main

import (
	_ "cmdb/models"
	_ "cmdb/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
