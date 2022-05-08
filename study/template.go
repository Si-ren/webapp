package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Template struct {
	var1 int
	var2 string
	var3 bool
}

//常用结构体
//Template
//常用函数:
//New 创建模板
//ParseFiles 指定文件模板
//ParseGlob 指定文件模板匹配格式
//Must 帮助函数,对模板创建结果进行验证,并返回模板对象指针
//常用方法:
//Parse 解析模板字符串
//ParseFiles 指定文件模板
//ParseGlob 指定文件模板匹配格式
//Execute 模板渲染
//ExecuteTemplate 指定模板执行模板渲染
//Funcs 指定自定义函数字典
//Clone 克隆模板进行模板复用

//创建处理器函数
func testTemplateIf(w http.ResponseWriter, r *http.Request) {
	////解析模板
	//t, _ := templates.ParseFiles("templates.html")
	////执行嵌入操作
	//t.Execute(w, "hello templates")

	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("if.html"))
	//解析指定的template.html
	T.ExecuteTemplate(w, "if.html", 15 > 16)
}

//创建处理器函数
func testTemplateRange(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("range.html"))
	var temps []*Template
	temp := &Template{
		var1: 0,
		var2: "string",
		var3: false,
	}
	temps = append(temps, temp)
	//解析指定的template.html
	T.ExecuteTemplate(w, "range.html", temps)
}

//创建处理器函数处理Map
func testTemplateMap(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("range.html"))

	var Maps map[string]string
	Maps = make(map[string]string)
	Maps["1"] = "one"
	Maps["2"] = "two"
	Maps["3"] = "three"
	//解析指定的template.html
	T.ExecuteTemplate(w, "range.html", Maps)
}

//创建处理器函数
func testTemplateWith(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("with.html"))

	T.ExecuteTemplate(w, "with.html", "Siri")
}

func testTemplateTemplate(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("templates.html", "if.html"))

	T.Execute(w, "Siri")
}

func testTemplateDefine(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("define1.html", "define2.html"))
	//模板这儿直接写定义的模板名
	T.ExecuteTemplate(w, "define1", "Siri")
}

//template使用自定义函数
//如果模板中要使用template,那么两个模板文件都需要解析
func testTemplateFunc(w http.ResponseWriter, r *http.Request) {
	tplFunc := template.FuncMap{
		"upper": strings.ToUpper,
		"title": func(text string) string {
			if len(text) == 0 {
				return ""
			} else if len(text) == 1 {
				return strings.ToUpper(text)
			}
			return strings.ToUpper(text[:1]) + text[2:]
		},
	}
	T := template.Must(template.New("tpl").Funcs(tplFunc).ParseFiles("templates.html", "define1.html"))
	err := T.ExecuteTemplate(w, "templates.html", "abcdefg")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/testTemplateIf", testTemplateIf)
	http.HandleFunc("/testTemplateRange", testTemplateRange)
	http.HandleFunc("/testTemplateMap", testTemplateMap)
	http.HandleFunc("/testTemplateWith", testTemplateWith)
	http.HandleFunc("/testTemplateTemplate", testTemplateTemplate)
	http.HandleFunc("/testTemplateDefine", testTemplateDefine)
	http.HandleFunc("/testTemplateFunc", testTemplateFunc)
	http.ListenAndServe(":8080", nil)
}
