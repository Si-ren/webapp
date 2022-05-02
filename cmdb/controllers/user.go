package controllers

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type UserController struct {
	Authentication
}

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
