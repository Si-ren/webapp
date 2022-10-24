package routers

import (
	"cmdb/controllers"
	v1 "cmdb/controllers/api/v1"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})

	beego.ErrorController(&controllers.ErrorController{})

	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.TargetController{})

	beego.AutoRouter(&v1.PrometheusController{}) //http://localhost:8080/prometheus/register
	v1 := beego.NewNamespace("/v1", beego.NSAutoRouter(&v1.PrometheusController{}))
	beego.AddNamespace(v1) //http://localhost:8080/v1/prometheus/register
}
