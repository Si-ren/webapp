package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	//logs.SetLogger("file", `{"filename":"test.log"}`)
	logs.SetLogger("console")
	//打印行数
	logs.EnableFuncCallDepth(true)
	//只有日志级别比info高的才会打印
	logs.SetLevel(logs.LevelDebug)
	logs.Error("错误日志")
	logs.Warning("警告日志")
	logs.Informational("普通日志")
	logs.Debug("调试日志")

}
