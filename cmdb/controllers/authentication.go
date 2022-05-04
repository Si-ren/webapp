package controllers

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type Authentication struct {
	beego.Controller
}

func (c *Authentication) Prepare() {
	sessionValue := c.GetSession("user")
	if sessionValue != nil {
		fmt.Println("")
		if ID, ok := sessionValue.(int); ok {
			if user := models.GetUserByID(ID); user != nil {
				c.Data["loginUser"] = user
				return
			}
		}
	}

	fmt.Println("AuthContoller Prepare")
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
	c.StopRun()
}
