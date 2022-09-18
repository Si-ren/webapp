package auth

import (
	"context"
	"github.com/sirupsen/logrus"
	authpb "grpcProtobuf/auth/api/v1"
)

type Service struct {
}

func (*Service) Login(c context.Context, req *authpb.LoginRequest) (loginResponse *authpb.LoginResponse, err error) {
	logrus.Info("This is auth rpc login")
	return &authpb.LoginResponse{
		AccessToken: "token-qweasdzxc",
		ExpiresIn:   2,
	}, nil
}
