package routers

import (
	"cmdb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})
}
