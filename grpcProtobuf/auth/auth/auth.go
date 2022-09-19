package auth

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "grpcProtobuf/auth/api/v1"
)

type Service struct {
	OAuthAuthentication OAuthAuthentication
}

type OAuthAuthentication interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (loginResponse *authpb.LoginResponse, err error) {
	logrus.Info("This is auth rpc login")
	OAuth, err := s.OAuthAuthentication.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve OAuth : %v", err)
	}

	return &authpb.LoginResponse{
		AccessToken: "token-qweasdzxc",
		ExpiresIn:   2,
	}, nil
}
