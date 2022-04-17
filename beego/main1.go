package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//***beego项目启动文件必须为main.go

type MyController struct {
	beego.Controller
}

func (c *MyController) GetMethod() {
	c.Ctx.WriteString("this is Get Method")

}
func (c *MyController) PostMethod() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(fmt.Sprintf("this is Post Method, ID is %s", id))

}
func (c *MyController) DeleteMethod() {
	c.Ctx.WriteString("this is Delete Method")

}

//提交数据方式:
//Get ?queryparam
//Post ?queryparam
//		request body
//		content-type: application/json
//					application/x-www-form-urlencoded
//					multipart/form-data

func (c *MyController) QueryParams() {
	//方式一
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.Form) //output: map[a:[1] b:[2]]

	//方式二
	fmt.Println(c.Ctx.Input.Query("a")) //output: 1

	//方式三,把值绑定到变量上去
	var a string
	c.Ctx.Input.Bind(&a, "a")
	fmt.Println(a) //output: 1

	//方式四
	fmt.Println(c.GetString("b")) //output: 2

	//方法五 ,解析json
	//如果不用反射,那么字段名称要和url中的键就要对应上,例如 /?A=1&B=2
	type InputForm struct {
		A string `form:"a"`
		B string `form:"b"`
	}
	var inputform InputForm
	fmt.Println(c.ParseForm(&inputform))
	fmt.Println(inputform)

	//方式六
	fmt.Println(c.Input())
	c.Ctx.WriteString("Success connect")
}

//获取request body中的字段
func (c *MyController) Form() {
	//方式一,拿取queryparam数据和form数据
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.Form)

	//方式二,仅拿取post method中的form数据
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.PostForm)

	//方式三,直接获取
	fmt.Println(c.GetString("aaa"))
	c.Ctx.WriteString("Success connect")

	//方式四,五,六 与 queryparams一样使用
}

//获取文件
func (c *MyController) File() {
	//1.  Request
	//2.  c.GetFile("image")
	c.SaveToFile("image", "./upload/UploadFile.txt")

	c.Ctx.WriteString("Success connect")

}

//获取json
//获取json要开启配置文件CopyRequestBody
func (c *MyController) Json() {
	var m map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	fmt.Println("json", m)
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	c.Ctx.WriteString("Success connect")

}

//获取cookie
func (c *MyController) Cookie() {
	//方法一
	cookie, err := c.Ctx.Request.Cookie("name")
	fmt.Println(cookie, err)
	//方法二
	fmt.Println(c.Ctx.Input.Cookie("name"))
	//方式三
	fmt.Println(c.Ctx.GetCookie("name"))

	c.Ctx.SetCookie("name", "bbbbbbbbb")

	//获取加密的cookie,此加密方式为sha256
	fmt.Println(c.Ctx.GetSecureCookie("passwd", "value"))
	c.Ctx.SetSecureCookie("passwd", "value", "aaaaaaa")
	c.Ctx.WriteString("Success connect")
}

//返回字符串
func (c *MyController) Output() {
	c.Ctx.WriteString("WriteString\n")
	c.Ctx.Output.Body([]byte("Output.Body\n"))
	c.Ctx.Output.Context.WriteString("OutPut.Context.WriteString\n")
}

//返回模板,模板一定要在views目录下,必须符合mvc
func (c *MyController) Tpl() {
	c.TplName = "output.html"
}

//返回JSON
func (c *MyController) RespondJson() {
	c.Data["json"] = map[string]string{"name": "siri", "wage": "99999"}
	c.ServeJSON()
}

//返回XML
func (c *MyController) RespondXML() {
	c.Data["xml"] = struct {
		XMLName xml.Name `xml:"root"`
		Name    string   `xml:"name"`
		Addr    string   `xml:"addr"`
	}{Name: "siri", Addr: "ShangHai"}
	c.ServeXML()
}

//返回YAML
func (c *MyController) RespondYAML() {
	c.Data["yaml"] = map[string]string{"a": "xx", "b": "yyy"}
	c.ServeYAML()
}

//redirect,不能与Redirect重名,因为其中已有Redirect方法
func (c *MyController) ReDirect() {
	c.Redirect("https://www.baidu.com", 302)
}

//停止
func (c *MyController) StopRun() {
	c.StopRun()
}

//错误处理
func (c *MyController) Ab() {
	c.Abort("404")
}

//获取header信息
func (c *MyController) Header() {
	fmt.Println(c.Ctx.Request.Method)
	fmt.Println(c.Ctx.Request.URL)
	fmt.Println(c.Ctx.Request.Header)
	//判断是否是Get请求
	fmt.Println(c.Ctx.Input.IsGet())
	c.Ctx.WriteString("Success connect")

}
func main() {
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	//自定义路由,并且取出id
	beego.Post("/id/:id(\\w+)/", func(c *context.Context) {
		//name := c.Input.Query("name")
		id := c.Input.Param(":id")
		c.WriteString(id)

		fmt.Println(id)
		//c.Output.Context.WriteString(fmt.Sprintf("post test %s", name))
	})
	//method对应自定义方法,对应关系用";"分隔,method用","分隔,
	beego.Router("/MyController/?:id", &MyController{}, "get:GetMethod;post:PostMethod;delete:DeleteMethod")

	//自动路由
	//url => 控制 controller/action
	// /My/GetMethod  => 对应的是MyController下的GetMethod方法
	beego.AutoRouter(&MyController{})
	beego.Run()
}
