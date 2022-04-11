package main

import (
	"fmt"
	"net/http"
)

//Cookie作用: 1.广告推荐  2.免登录
//Cookie缺点: 如果cookie很多,那么就增加了数据传输量.浏览器对Cookie数量有限制,无法用cookie保存过多信息
//Session: 在服务器端保存一些用户数据,与一个Cookie相关联.
//Session原理:
//1)第一次向服务 器发送请求时创建Session, 给它设置-个全球唯一-的ID(可以通过UUID生成).
//2)创建一 个Cookie,将Cookie的Value设置为Session 的ID值，并将Cookie发送给浏览器.
//3)以后再发送请求浏览器就会携带着该Cookie.
//4)服务 器获取Cookie并根据它的Value值找到服务器中对应的Session,也就知道了请求是那个用户发的.

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie01 := http.Cookie{
		Name:     "user01",
		Value:    "admin01",
		HttpOnly: true,
		MaxAge:   10,
	}
	cookie02 := http.Cookie{
		Name:     "user02",
		Value:    "admin02",
		HttpOnly: true,
		MaxAge:   20,
	}
	cookie03 := http.Cookie{
		Name:     "user03",
		Value:    "admin03",
		HttpOnly: true,
		MaxAge:   30,
	}
	//将cookie发送给浏览器
	w.Header().Set("Set-Cookie", cookie01.String())
	//添加cookie
	w.Header().Add("Set-Cookie", cookie02.String())
	//直接调用http的SetCookie方法设置cookie
	http.SetCookie(w, &cookie03)
	//结果为
	//Set-Cookie: user01=admin01; HttpOnly
	//Set-Cookie: user02=admin02; HttpOnly
	//Set-Cookie: user03=admin03; HttpOnly
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	//获取全部cookie
	cookie := r.Header["Cookie"]
	fmt.Println("Get Cookie: ", cookie)
	//获取指定cookie
	cookie01, _ := r.Cookie("user01")
	cookie02, _ := r.Cookie("user02")
	cookie03, _ := r.Cookie("user03")
	fmt.Println("Get Cookie01: ", cookie01)
	fmt.Println("Get Cookie02: ", cookie02)
	fmt.Println("Get Cookie03: ", cookie03)

	fmt.Println("Content-Length: ", r.Header["Content-Length"])

}

func main() {
	http.HandleFunc("/Cookie", SetCookie)
	http.HandleFunc("/GetCookie", GetCookie)
	http.ListenAndServe(":8080", nil)
}
