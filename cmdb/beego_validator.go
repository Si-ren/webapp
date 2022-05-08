package main

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func main() {
	valid := &validation.Validation{}
	text := ""
	//验证
	valid.Required(text, "a.b.c")
	valid.Alpha("123456", "alpha.alpha.alpha").Message("输入大小写只能是大小写英文字母")
	valid.MaxSize("123456", 3, "alpha.alpha.alpha")

	//获取验证结果
	fmt.Println(valid.HasErrors())
	if valid.HasErrors() {
		fmt.Println(valid.Errors)
		fmt.Println(valid.ErrorsMap)
	}
}
