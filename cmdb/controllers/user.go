package controllers

import (
	"cmdb/forms"
	"cmdb/models"
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
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		//无session信息,未登录
		//sessionUser断言 ->int
		//user状态 -> 禁用/已离职
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		return
	}
	query := c.Ctx.Input.Query("query")
	users, err := models.QueryUser(query)
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
			models.ModifyUserByForm(form)
			c.Redirect(beego.URLFor("UserController.GetUser"), http.StatusFound)
		}
	} else if ID, err := c.GetInt("id"); err == nil {
		fmt.Println("Get ID: ", ID)
		if user := models.GetUserByID(ID); user != nil {
			form.ID = user.ID
			form.Name = user.Name
			form.Password = user.Password
		}
	}
	c.Data["form"] = form
	c.TplName = "user/modify.html"
}

func (c *UserController) Delete() {
	if id, err := c.GetInt("id"); err == nil && c.LoginUser.ID != id {
		models.DeleteUserByID(id)
	}
	c.Redirect(beego.URLFor("UserController.GetUser"), http.StatusFound)

}
