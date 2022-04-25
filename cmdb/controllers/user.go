package controllers

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUser() {
	query := c.Ctx.Input.Query("query")
	users, err := models.QueryUser(query)
	if err != nil {
		fmt.Println(users, err)
	}
	c.Data["users"] = users
	c.TplName = "user/query.html"
}
