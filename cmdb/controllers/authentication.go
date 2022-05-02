package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type Authentication struct {
	beego.Controller
}

func (c *Authentication) Prepare() {
	if c.GetSession("user") == nil {
		fmt.Println("AuthContoller Prepare")
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		c.StopRun()
	}
}
