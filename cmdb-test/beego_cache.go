package main

import (
	"fmt"
	cache2 "github.com/astaxie/beego/cache"
)

func main() {

	cache, _ := cache2.NewCache("memory", `{"interval":60}`)
	//cache, _ := cache2.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

	fmt.Println(cache.Get("name"))
	cache.Put("name", "siri", 10)
	fmt.Println(cache.Get("name"))
	fmt.Println(cache.IsExist("name"))
	//time.Sleep(12 * time.Second)
	cache.Delete("name")
	fmt.Println(cache.Get("name"))

	user := struct {
		ID   int
		Name string
	}{1, "siri"}
	cache.Put("user", user, 20)
	fmt.Println(cache.Get("user"))

}
