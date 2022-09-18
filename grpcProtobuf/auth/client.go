package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authpb "grpcProtobuf/auth/api/v1"
)

func main() {
	con, _ := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := authpb.NewAuthServiceClient(con)
	r, err := client.Login(context.Background(), &authpb.LoginRequest{Code: "siri"})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(r)
}
