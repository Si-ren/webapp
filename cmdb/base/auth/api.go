package auth

import (
	"cmdb/base/response"
	"fmt"
	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) Prepare() {
	c.EnableXSRF = false
	token := fmt.Sprintf("Token %s", beego.AppConfig.DefaultString("api::token", ""))
	fmt.Printf("%#v\n", token)
	headerToken := c.Ctx.Input.Header("Authorization")
	//headerToken := c.GetString("Authorization")
	fmt.Printf("APIController c.Ctx.Input.Header: %#v\n", headerToken)
	if token != headerToken {
		c.Data["json"] = response.Unauthorzation
		c.ServeJSON()
		c.StopRun()
	}
}

func (c *APIController) Render() error {
	c.ServeJSON()
	return nil
}
