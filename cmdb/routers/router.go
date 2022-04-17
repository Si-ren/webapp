package routers

import (
	"github.com/astaxie/beego"
	"webapp/cmdb/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
}
