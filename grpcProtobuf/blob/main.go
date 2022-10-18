package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpcProtobuf/auth/dao"
	blobpb "grpcProtobuf/blob/api/v1"
	"grpcProtobuf/blob/blob"
	"grpcProtobuf/common/server"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	c := context.Background()
	mgc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://siri:Wxr0910+1s@122.51.16.26:17018/grpcProtobuf"))
	if err != nil {
		log.Fatal("cannot connect mongodb", err)
	}
	col := dao.NewMongodb(mgc.Database("grpcProtobuf"))

	logrus.Fatal(server.RunGpcServer(&server.ConfigServer{
		Name:    "BLOB",
		Network: "tcp",
		Address: "8085",
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Storage: nil,
				Mongodb: dao.Mongodb{col},
			})
		},
		Logger: log,
	}))

}
