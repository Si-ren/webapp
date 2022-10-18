package main

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	authpb "grpcProtobuf/auth/api/v1"
	"grpcProtobuf/auth/auth"
	"grpcProtobuf/auth/dao"
	"grpcProtobuf/auth/utils"
	"grpcProtobuf/common/server"
	"io"
	"os"
	"time"
)

func main() {
	log := logrus.New()
	pkFile, err := os.Open("private.key")
	if err != nil {
		logrus.Panic("OS Open privateKey file err: ", err)
	}
	pkBytes, err := io.ReadAll(pkFile)
	prik, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logrus.Panic("IO Read privateKey file err: ", err)
	}
	logrus.Info("Start auth server")

	c := context.Background()
	mgc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://siri:Wxr0910+1s@122.51.16.26:17018/grpcProtobuf"))
	col := dao.NewMongodb(mgc.Database("grpcProtobuf"))

	logrus.Fatal(server.RunGpcServer(&server.ConfigServer{
		Name:    "auth",
		Network: "tcp",
		Address: ":8083",
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				//OAuthAuthentication: nil,
				Mongodb:       col,
				TokenGenerate: utils.NewJWTToken("auth", prik),
				TokenExpire:   2 * time.Hour,
			})
		},
		Logger: log,
	}))
}
