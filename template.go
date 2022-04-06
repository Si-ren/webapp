package main

import (
	"html/template"
	"net/http"
)

type Template struct {
	var1 int
	var2 string
	var3 bool
}

//创建处理器函数
func testTemplateIf(w http.ResponseWriter, r *http.Request) {
	////解析模板
	//t, _ := template.ParseFiles("template.html")
	////执行嵌入操作
	//t.Execute(w, "hello template")

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

//创建处理器函数
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
	T := template.Must(template.ParseFiles("template.html", "if.html"))

	T.Execute(w, "Siri")
}

func testTemplateDefine(w http.ResponseWriter, r *http.Request) {
	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("define1.html", "define2.html"))
	//模板这儿直接写定义的模板名
	T.ExecuteTemplate(w, "define1", "Siri")
}

func main() {
	http.HandleFunc("/testTemplateIf", testTemplateIf)
	http.HandleFunc("/testTemplateRange", testTemplateRange)
	http.HandleFunc("/testTemplateMap", testTemplateMap)
	http.HandleFunc("/testTemplateWith", testTemplateWith)
	http.HandleFunc("/testTemplateTemplate", testTemplateTemplate)
	http.HandleFunc("/testTemplateDefine", testTemplateDefine)
	http.ListenAndServe(":8080", nil)
}
