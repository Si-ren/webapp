package main

import (
	"fmt"
	"net/http"
	"time"
)

func muxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建多路复用器处理请求", r.URL.Path)
}

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world", r.URL.Path)
}

type MyHandler struct {
}

//使用自定义server实例,方法名必须为 ServeHTTP
func (myhandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "MyHandle test", r.URL.Path)
}

func main() {

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	////自定义多路复用器
	//mux := http.NewServeMux()
	////使用自己的多路复用器
	//mux.HandleFunc("/", muxHandler)
	//http.ListenAndServe(":8080", mux)

	//使用自定义server实例
	//自定义多路复用器
	mux := http.NewServeMux()
	//使用自己的多路复用器
	mux.HandleFunc("/", muxHandler)
	mux.Handle("/test", &MyHandler{})
	//创建http.Server结构体,里面详细配置
	server := &http.Server{
		//Addr写IP+端口，如果不写ip,端口格式 --> :80
		Addr:    "127.0.0.1:8080",
		Handler: mux,
		//TLSConfig:         nil,
		//ReadTimeout: 2 * time.Second,
		//ReadHeaderTimeout: 0,
		WriteTimeout: 2 * time.Second,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}

	server.ListenAndServe()
}
