package services

import (
	"cmdb/conf"
	"cmdb/services/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	Host user.Service
	//把服务注册到svcs,ginSvcs中
	//让这俩管理所有service的Config和注册路由
	implSvcs = map[string]ImplService{}
	ginSvcs  = map[string]GinService{}
	grpcSvcs = map[string]GrpcService{}
)

type ImplService interface {
	Config() error
	Name() string
}

// ImplRegister 服务绑定
func ImplRegister(svc ImplService) {
	if _, ok := implSvcs[svc.Name()]; ok {
		panic(fmt.Sprintf("ImplService %s has registered", svc.Name()))
	}
	implSvcs[svc.Name()] = svc
	if v, ok := svc.(user.Service); ok {
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

func InitGinSvc(r gin.IRouter) {
	for k, v := range ginSvcs {
		if err := v.Configure(); err != nil {
			conf.Log.Panicf("%s router handler Configure  err : %s", k, err)
		}
		v.Registry(r)
		conf.Log.Infof("%s router handler  register router  complete", k)
	}
}

// GrpcService 定义grpc service
type GrpcService interface {
	Registry(g *grpc.Server)
	Config() error
	Name() string
}

// GrpcRegister 在grpc服务init时 注册到ioc中
func GrpcRegister(svc GrpcService) {
	_, ok := grpcSvcs[svc.Name()]
	if !ok {
		panic(fmt.Sprintf("GrpcService %s has registered", svc.Name()))
	}
	grpcSvcs[svc.Name()] = svc
}

// InitGRPCSvc 初始化所有grpc服务
func InitGRPCSvc(server *grpc.Server) error {
	for name, svc := range grpcSvcs {
		if err := svc.Config(); err != nil {
			return fmt.Errorf("init grpc service %s err: %s", name, err.Error())
		}
		svc.Registry(server)
	}
	return nil
}

// GetGrpcSvcs 获取所有注册的grpc服务
func GetGrpcSvcs() (names []string) {
	for k := range grpcSvcs {
		names = append(names, k)
	}
	return names
}
