package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	classpb "grpcProtobuf/class/api/v1"
	"grpcProtobuf/class/class"
	"grpcProtobuf/common/auth"
	"net"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Info("Start class service")
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		logrus.Panic("Class listen error :", err)
	}

	in, err := auth.Interceptor("common/auth/public.key")
	if err != nil {
		logrus.Fatal("Cant create Interceptor :", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(in))

	classpb.RegisterClassServiceServer(s, &class.Service{})
	log.Fatalln(s.Serve(lis))

}
