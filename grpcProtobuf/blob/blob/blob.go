package blob

import (
	"context"
	"github.com/sirupsen/logrus"
	blobpb "grpcProtobuf/blob/api/v1"
	"grpcProtobuf/blob/dao"
	"grpcProtobuf/common/id"
	"io"
	"net/http"
	"time"
)

type Storage interface {
	SignURL(c context.Context, method string, path string, timeout time.Duration) (string, error)
	Get(c context.Context, path string) (io.ReadCloser, error)
}

type Service struct {
	Storage Storage
	Mongodb *dao.Mongodb
	log     *logrus.Logger
}

func (s *Service) CreateBolb(ctx context.Context, rep *blobpb.CreateBolbRequest) (*blobpb.CreateBolbResponse, error) {
	aid := id.AccountID(rep.AccountId)
	blobRes, err := s.Mongodb.CreateBolb(ctx, aid)
	if err != nil {
		s.log.Error("Blob create blob err :", err)
		return nil, err
	}
	u, err := s.Storage.SignURL(ctx, http.MethodPut, blobRes.Path, secToDuration(10))
	if err != nil {
		s.log.Error("Blob CreateBolb  SignURL err :", err)
		return nil, err
	}

	return &blobpb.CreateBolbResponse{
		Id:      blobRes.AccountID,
		BolbUrl: u,
	}, nil
}

func (s *Service) GetBolbURL(c context.Context, req *blobpb.GetBolbUrlRequest) (*blobpb.GetBolbUrlResponse, error) {
	aid := id.AccountID(req.AccountId)
	s.Mongodb.Get
	return nil, nil
}

func secToDuration(sec int32) time.Duration {
	return time.Duration(sec) * time.Second
}
