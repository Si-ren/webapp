package main

import (
	"cmdb/cmds"
	_ "cmdb/routers"
	_ "github.com/go-sql-driver/mysql" //需要导入数据库所需 driver
	"time"
)

func main() {
	//models.CacheInit("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	//beego.Run()
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	cmds.Execute()
}
