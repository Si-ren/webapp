package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"webapp/cmdb/base/error"
	"webapp/cmdb/forms"
	"webapp/cmdb/models"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	form := &forms.LoginForm{}
	error := error.New()
	if c.Ctx.Input.IsPost() {
		form.Name = c.Ctx.Input.Query("name")
		form.Password = c.Ctx.Input.Query("password")
		fmt.Println(form)
		user := models.GetUserByName(form.Name)
		fmt.Println(user)
		if user == nil {
			//用户不存在
			error.AddError("default", "用户名或密码不正确")
		} else if user.ValidPassword(form.Password) {
			//用户密码正确
			c.Redirect("/home/index", http.StatusFound)

		} else {
			//用户密码不正确
			error.AddError("default", "用户名或密码不正确")
		}
	}
	fmt.Println(error)
	c.Data["form"] = form
	c.Data["error"] = error
	c.TplName = "auth/login.html"
}
