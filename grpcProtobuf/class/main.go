package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	classpb "grpcProtobuf/class/api/v1"
	"grpcProtobuf/class/class"
	"net"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Info("Start class service")
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		logrus.Panic("Class listen error :", err)
	}
	classpb.RegisterClassServiceServer(s, &class.Service{})
	log.Fatalln(s.Serve(lis))
}
