package auth

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "grpcProtobuf/auth/api/v1"
	"grpcProtobuf/auth/dao"
	"time"
)

type OAuthAuthentication interface {
	Resolve(code string) (string, error)
}

type TokenGenerate interface {
	TokenGenerate(string, time.Duration) (string, error)
}

type Service struct {
	// OAuthAuthentication OAuthAuthentication
	Mongodb       *dao.Mongodb
	TokenGenerate TokenGenerate
	TokenExpire   time.Duration
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (loginResponse *authpb.LoginResponse, err error) {
	//logrus.Info("This is auth rpc login")
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
	token, err := s.TokenGenerate.TokenGenerate(accountID, s.TokenExpire)
	if err != nil {
		logrus.Error("AccountID generate token err:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}

//func (s *Service) Resolve(code string) (string, error) {
//	return code, nil
//}
