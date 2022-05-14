package controllers

import (
	"cmdb/forms"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type UserController struct {
	Authentication
}

//func (c *UserController) Prepare() {
//	c.Authentication.Prepare()
//	c.Data["nav"] = "user"
//}

func (c *UserController) GetUser() {
	//flash读取 flash存储的值, 读取完会把c.data里面的值删掉
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		//无session信息,未登录
		//sessionUser断言 ->int
		//user状态 -> 禁用/已离职
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		return
	}
	query := c.Ctx.Input.Query("query")
	users, err := services.UserService.Query(query)
	if err != nil {
		fmt.Println(users, err)
	}
	c.Data["users"] = users
	c.TplName = "user/query.html"
}

func (c *UserController) Modify() {
	form := &forms.UserModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			fmt.Println(form)
			services.UserService.ModifyByForm(form)
			//flash.set会同时把 key value 放入data中,即c.Data["key"] = value
			flash := beego.NewFlash()
			//flash.Notice("修改用户信息成功")等同于flash.Set("notice","修改用户信息成功")
			flash.Set("notice", "修改用户信息成功")
			//flash.Success("success") 等同于flash.Set("success", "success")
			//flash.Error("error") 等同于flash.Set("error", "error")
			//flash.Warning("warning") 等同于flash.Set("warning", "warning")
			flash.Store(&c.Controller)
			c.Redirect(beego.URLFor("UserController.GetUser"), http.StatusFound)
		}
	} else if ID, err := c.GetInt("id"); err == nil {
		fmt.Println("Get ID: ", ID)
		if user := services.UserService.GetByID(ID); user != nil {
			form.ID = user.ID
			form.Name = user.Name
			form.Password = user.Password
		}
	}
	c.Data["form"] = form
	//直接生成xsrf的token,在界面上拼接,试了下没用卧槽
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/modify.html"
	c.Layout = "layout/layout.html"
}

func (c *UserController) Delete() {
	if id, err := c.GetInt("id"); err == nil && c.LoginUser.ID != id {
		services.UserService.DeleteByID(id)
	}
	c.Redirect(beego.URLFor("UserController.GetUser"), http.StatusFound)

}
