package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func testTemplate(w http.ResponseWriter, r *http.Request) {
	////解析模板
	//t, _ := template.ParseFiles("template.html")
	////执行嵌入操作
	//t.Execute(w, "hello template")

	//使用must函数来对err进行处理, 如果err非nil, 那么must函数会产生panic
	T := template.Must(template.ParseFiles("template.html"))
	//解析指定的template.html
	T.ExecuteTemplate(w, "template.html", "hello ExecuteTemplate")
}
func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
