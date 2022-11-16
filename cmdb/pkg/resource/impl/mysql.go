package impl

import (
	"cmdb/conf"
	"cmdb/pkg"
	resource "cmdb/pkg/resource/api/v1"

	"gorm.io/gorm"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

var (
	svc                     = &service{}
	_       pkg.GrpcService = (*service)(nil)
	AppName                 = "resource"
)

// 定义grpc服务
type service struct {
	db  *gorm.DB
	log *logrus.Logger
	// 需要这个是因为这个结构体包含所有接口 函数
	resource.UnimplementedServiceServer
}

// Registry  resource服务注册
func (s *service) Registry(g *grpc.Server) {
	resource.RegisterServiceServer(g, s)
}

func (s *service) Config() error {
	db, err := conf.Configure().MySQL.GetDB()
	if err != nil {
		return err
	}
	//初始化 resource 服务
	s.log = conf.Log
	s.db = db
	//为什么在这里初始化数据表
	//因为在init时，svc的db还没初始化好
	s.db.AutoMigrate(&resource.Resource{})
	s.db.AutoMigrate(&resource.Tag{})
	s.db.AutoMigrate(&resource.Base{})
	s.db.AutoMigrate(&resource.Information{})
	return nil
}

func (s *service) Name() string {
	return AppName
}

func init() {
	pkg.GrpcRegister(svc)
}
