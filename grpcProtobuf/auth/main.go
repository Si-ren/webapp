package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	authpb "grpcProtobuf/auth/api/v1"
	"grpcProtobuf/auth/auth"
	"net"
)

func main() {

	logrus.Info("Start auth server")
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		logrus.Panic("Auth listen error")
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{})
	logrus.Fatal(s.Serve(lis))
}
