package controllers

import (
	"cmdb/models"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type Authentication struct {
	beego.Controller
	LoginUser *models.User
}

func (c *Authentication) Prepare() {
	//这里是layout的功能,如果html中有用到数据,必须传值,没有值则传"",例:c.Data["subnav"] = ""
	controllerName, actionName := c.GetControllerAndAction()
	fmt.Println(controllerName, actionName)
	c.Data["nav"] = controllerName
	c.Data["subnav"] = ""
	//session功能
	sessionValue := c.GetSession("user")
	if sessionValue != nil {
		fmt.Println("")
		if ID, ok := sessionValue.(int); ok {
			if user := services.UserService.GetByID(ID); user != nil {
				c.LoginUser = user
				c.Data["loginUser"] = user
				return
			}
		}
	}

	fmt.Println("AuthContoller Prepare")
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
	c.StopRun()
}
