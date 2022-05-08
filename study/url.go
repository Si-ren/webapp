package main

import (
	"encoding/json"
	"fmt"
	"helo/study/model"
	"net/http"
)

//创建处理器函数
func handler01(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的Path是", r.URL.Path)
	fmt.Fprintln(w, "你发送的查询字符串是", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是: ", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性是: ", r.Header.Get("Accept-Encoding"))

	//	获取请求体中内容长度
	// ******注意!body只允许读取一次
	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求体中内容: ", string(body))

	//如果已经读取了body,那么如果再解析,那么会出现nil的情况
	//解析表单,调用r.Form之前必须执行的操作
	err := r.ParseForm()
	fmt.Println("r.ParseForm err: ", err)
	//fmt.Fprintln(w, "请求参数: ", r.Form)
	//fmt.Fprintln(w, "POST请求的参数: ", r.PostForm)
	fmt.Fprintln(w, "POST请求的参数的值: ", r.PostFormValue("username"))
	fmt.Fprintln(w, "URL中的user请求参数值: ", r.FormValue("user.go"))
}

//返回json
func TestJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{
		ID:       2,
		Username: "siri",
		Password: "pwd",
		Email:    "test@test.com",
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	w.Write(data)
}

//客户端重定向
func TestRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/test", handler01)
	http.HandleFunc("/TestJsonRes", TestJsonRes)
	http.HandleFunc("/baidu", TestRedirect)
	http.ListenAndServe(":8080", nil)
}
