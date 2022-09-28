package auth

import (
	"go/build"
	"google.golang.org/grpc"
)

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	return nil, nil

}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx build.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) {

}
