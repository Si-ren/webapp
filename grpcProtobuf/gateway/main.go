package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	authpb "grpcProtobuf/auth/api/v1"
	classpb "grpcProtobuf/class/api/v1"
	"net/http"
)

func main() {
	logrus.Info("GRPC Gateway start")
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true},
		UnmarshalOptions: protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true},
	}))
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, gwmux, ":8083", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logrus.Panic("Auth Service Register Gateway Failed")
		return
	}
	err = classpb.RegisterClassServiceHandlerFromEndpoint(c, gwmux, ":8084", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logrus.Panic("Class Service Register Gateway Failed")
		return
	}
	logrus.Fatal(http.ListenAndServe(":8082", gwmux))
}
