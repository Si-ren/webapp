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
	svc                 = &service{}
	_   pkg.GrpcService = (*service)(nil)
)

type service struct {
	db  *gorm.DB
	log *logrus.Logger
	resource.UnimplementedServiceServer
}

func (s *service) Registry(g *grpc.Server) {
	resource.RegisterServiceServer(g, s)
}

func (s *service) Config() error {
	db, err := conf.Configure().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = conf.Log
	s.db = db
	s.db.AutoMigrate(&resource.Resource{})
	s.db.AutoMigrate(&resource.Tag{})
	s.db.AutoMigrate(&resource.Base{})
	s.db.AutoMigrate(&resource.Information{})
	return nil
}

func (s *service) Name() string {
	return "resource"
}

func init() {
	pkg.GrpcRegister(svc)
}
