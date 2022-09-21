package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	authpb "grpcProtobuf/auth/api/v1"
	"grpcProtobuf/auth/auth"
	"grpcProtobuf/auth/dao"
	"net"
)

func main() {

	logrus.Info("Start auth server")
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		logrus.Panic("Auth listen error")
	}
	s := grpc.NewServer()
	c := context.Background()
	mgc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://192.168.179.128:27017"))
	col := dao.NewMongodb(mgc.Database("grpcProtobuf"))
	authpb.RegisterAuthServiceServer(s, &auth.Service{OAuthAuthentication: nil, Mongodb: col})
	logrus.Fatal(s.Serve(lis))
}
