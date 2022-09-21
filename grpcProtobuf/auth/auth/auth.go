package auth

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "grpcProtobuf/auth/api/v1"
	"grpcProtobuf/auth/dao"
)

type Service struct {
	OAuthAuthentication OAuthAuthentication
	Mongodb             *dao.Mongodb
}

type OAuthAuthentication interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (loginResponse *authpb.LoginResponse, err error) {
	logrus.Info("This is auth rpc login")
	//OAuth, err := s.OAuthAuthentication.Resolve(req.Code)
	//if err != nil {
	//	return nil, status.Errorf(codes.Unavailable, "cannot resolve OAuth : %v", err)
	//}
	OAuth := req.Code
	logrus.Info(req.Code)
	accountID, err := s.Mongodb.FindID(c, OAuth)
	if err != nil {
		logrus.Error("cant resolve account id :", err)
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		AccessToken: accountID,
		ExpiresIn:   2,
	}, nil
}

func (s *Service) Resolve(code string) (string, error) {
	return code, nil
}
