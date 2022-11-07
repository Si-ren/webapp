package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderAccessKey = "client-id"
	ClientHeaderSecretKey = "client-secret"
)

func NewAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&GRPCAuth{}).Auth
}

type GRPCAuth struct {
}

func (g *GRPCAuth) Auth(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	//从metadata读取凭证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context: %s", err)
	}
	clientID, clientSecret := g.getClientCredentialsFromMeta(md)
	if err = g.validateServiceCredential(clientID, clientSecret); err != nil {
		return nil, err
	}

	return handler(ctx, req)

}

func (g *GRPCAuth) getClientCredentialsFromMeta(md metadata.MD) (clientID, clientSecret string) {
	cakList := md[ClientHeaderAccessKey]
	if len(cakList) > 0 {
		clientID = cakList[0]
	}
	cskList := md[ClientHeaderSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}
	return clientID, clientSecret
}

func (g GRPCAuth) validateServiceCredential(clientID, clientSecret string) (err error) {
	if !(clientID == "admin" && clientSecret == "admin") {
		return status.Errorf(codes.Unauthenticated, "clientID, clientSecret error")
	}
	return nil
}
