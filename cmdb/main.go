package main

import (
	"cmdb/cmds"
	_ "cmdb/routers"
	_ "github.com/go-sql-driver/mysql" //需要导入数据库所需 driver
)

func main() {
	//models.CacheInit("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	//beego.Run()
	cmds.Execute()
}
