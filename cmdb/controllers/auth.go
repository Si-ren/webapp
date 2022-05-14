package controllers

import (
	"cmdb/base/error"
	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	models.Cache.Incr("LoginNum")
	if c.GetSession("user") != nil {
		c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		return
	}
	form := &forms.LoginForm{}
	error := error.New()
	if c.Ctx.Input.IsPost() {
		form.Name = c.Ctx.Input.Query("name")
		form.Password = c.Ctx.Input.Query("password")
		fmt.Println("Get Form:", form)
		user, err := services.UserService.GetByName(form.Name)
		if err != nil {
			fmt.Println(err)
		}
		if user == nil {
			//用户不存在
			error.AddError("default", "用户不存在")
		} else if user.ValidPassword(form.Password) {
			//用户密码正确
			//记录用户状态(session,记录服务器)

			c.SetSession("user", user.ID)
			c.Redirect(beego.URLFor("UserController.GetUser"), http.StatusFound)
			//c.Redirect("/user/getuser", http.StatusFound)

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

func (c *AuthController) Register() {
	error := error.New()
	user := models.User{
		StaffID:    "",
		Name:       c.GetString("name"),
		NickName:   c.GetString("nickname"),
		Password:   c.GetString("password"),
		Tel:        c.GetString("telephone"),
		Addr:       c.GetString("address"),
		Email:      c.GetString("email"),
		Department: c.GetString("department"),
	}
	user.Gender, _ = c.GetInt("gender")
	fmt.Println(user)
	if c.Ctx.Input.IsPost() {
		users, err := services.UserService.GetByName(user.Name)
		fmt.Println(err)
		fmt.Println(users)
		if err == nil {
			fmt.Println("用户名已经存在")
			error.AddError("Register error: ", "用户名已存在")
		} else {

			fmt.Println("")
			services.UserService.Create(&user)
		}
	}
	fmt.Println(error)
	c.Data["error"] = error
	c.TplName = "auth/register.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
}
