package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

//***beego项目启动文件必须为main.go

type TemplateController struct {
	beego.Controller
}

func (c *TemplateController) Tpl() {
	c.Data["name"] = "siri"
	c.Data["gender"] = "true"
	c.Data["socre"] = []float32{1, 2, 3, 4}
	c.Data["users"] = map[int]string{1: "siri", 2: "lsl"}
	c.Data["content"] = "abc.ABC"
	//后缀只支持html和tpl
	c.TplName = "templates.html"

}
func main() {
	//获取配置文件中配置信息
	fmt.Println(beego.AppConfig.String("MYSQL_HOST"))
	fmt.Println(beego.AppConfig.Int("MYSQL_PORT"))

	beego.AddFuncMap("lower", func(low string) string {
		return strings.ToLower(low)
	})
	beego.AutoRouter(&TemplateController{})
	beego.Run()
}
