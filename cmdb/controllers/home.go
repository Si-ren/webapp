package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

type HomeController struct {
	Authentication
}

func (c *HomeController) Index() {
	fmt.Println("Index")
	//session检查
	c.GetSession("user")
	c.TplName = "home/index.html"
}

func (c *HomeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	fmt.Println("Init: ", controllerName, actionName)
}

//func (c *HomeController) Prepare() {
//	c.Authentication.Prepare()
//	c.Data["nav"] = "home"
//}

func (c *HomeController) Test() {
	fmt.Println("Test")
	c.Ctx.WriteString("Test")
}

//如果有返回,那么不会执行render
func (c *HomeController) Render() error {
	fmt.Println("Render")
	return c.Controller.Render()
}

func (c *HomeController) Finish() {
	fmt.Println("Finish")
}
