package protocol

import (
	"cmdb/conf"
	"cmdb/pkg"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// GRPCService http服务
type GRPCService struct {
	g *grpc.Server
	l *logrus.Logger
	c *conf.Config
}

// NewHTTPService 构建函数
func NewGRPCService() *GRPCService {
	server := grpc.NewServer()
	return &GRPCService{
		g: server,
		l: conf.Log,
		c: conf.Configure(),
	}
}

func (s *GRPCService) Start() error {
	s.l.Info("Start grpc service")
	pkg.InitConfigSvc()
	if err := pkg.InitGRPCSvc(s.g); err != nil {
		return err
	}
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		return err
	}
	s.l.Infof("listen grpc tcp connect err: %s", s.c.App.GRPCAddr())
	if err := s.g.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}
		s.l.Error("start grpc service error, %s", err.Error())
		return err
	}

	return nil

}

func (s *GRPCService) Stop() {
	s.g.GracefulStop()
}
