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
	//初始化grpc server
	server := grpc.NewServer()
	return &GRPCService{
		g: server,
		l: conf.Log,
		c: conf.Configure(),
	}
}

// Start grpc服务启动
func (s *GRPCService) Start() error {
	s.l.Info("Start grpc service")

	pkg.InitConfigSvc()
	//初始化所有grpc服务
	if err := pkg.InitGRPCSvc(s.g); err != nil {
		return err
	}
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		return err
	}
	s.l.Infof("listen grpc tcp connect addr: %s", s.c.App.GRPCAddr())
	//启动服务
	if err := s.g.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}
		s.l.Error("start grpc service error, %s", err.Error())
		return err
	}

	return nil

}

// Stop grpc优雅关闭
func (s *GRPCService) Stop() {
	s.g.GracefulStop()
}
