package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"grpcProtobuf/common/auth"
	"net"
)

type ConfigServer struct {
	Name          string
	Network       string
	Address       string
	AuthPublicKey string
	RegisterFunc  func(server *grpc.Server)
	Logger        *logrus.Logger
}

func RunGpcServer(c *ConfigServer) error {
	lis, err := net.Listen(c.Network, c.Address)
	if err != nil {
		c.Logger.Fatal(c.Name, " server listen error : ", err)
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKey != "" {
		in, err := auth.Interceptor(c.AuthPublicKey)
		if err != nil {
			c.Logger.Fatal("Server ", c.Name, " auth Interceptor err: ", err)
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	c.Logger.Info("server started", c.Name, "addr", c.Address)
	return s.Serve(lis)
}
