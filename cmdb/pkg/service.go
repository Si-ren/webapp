package pkg

import (
	"cmdb/conf"
	"cmdb/pkg/host"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	Host host.Service
	//把服务注册到svcs,ginSvcs中
	//让这俩管理所有service的Config和注册路由
	implSvcs = map[string]ImplService{}
	ginSvcs  = map[string]GinService{}
)

type ImplService interface {
	Config() error
	Name() string
}

func ImplRegister(svc ImplService) {
	if _, ok := implSvcs[svc.Name()]; ok {
		panic(fmt.Sprintf("ImplService %s has registered", svc.Name()))
	}
	implSvcs[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		Host = v
	}
}

func InitConfigSvc() {
	for k, v := range implSvcs {
		if err := v.Config(); err != nil {
			conf.Log.Panic("%s service init err :", err)
		}
		conf.Log.Infof("%s service init complete", k)
	}
}

type GinService interface {
	Registry(r gin.IRouter)
	Name() string
	Configure() error
}

func GinRegister(svc GinService) {
	if _, ok := ginSvcs[svc.Name()]; ok {
		panic(fmt.Sprintf("GinService %s has registered", svc.Name()))
	}
	ginSvcs[svc.Name()] = svc
}

func GetGinSvcs() (names []string) {
	for k := range ginSvcs {
		names = append(names, k)
	}
	return names
}

func InitRouterSvc(r gin.IRouter) {
	for k, v := range ginSvcs {
		if err := v.Configure(); err != nil {
			conf.Log.Panicf("%s router handler Configure  err : %s", k, err)
		}
		v.Registry(r)
		conf.Log.Infof("%s router handler  register router  complete", k)
	}
}