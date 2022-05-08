package controllers

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type Authentication struct {
	beego.Controller
	LoginUser *models.User
}

func (c *Authentication) Prepare() {
	sessionValue := c.GetSession("user")
	controllerName, actionName := c.GetControllerAndAction()
	fmt.Println(controllerName, actionName)
	c.Data["nav"] = controllerName
	if sessionValue != nil {
		fmt.Println("")
		if ID, ok := sessionValue.(int); ok {
			if user := models.GetUserByID(ID); user != nil {
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
