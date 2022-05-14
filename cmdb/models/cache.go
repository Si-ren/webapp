package models

import "github.com/astaxie/beego/cache"

var Cache cache.Cache

func CacheInit(adapter, config string) {
	Cache, _ = cache.NewCache(adapter, config)

}
