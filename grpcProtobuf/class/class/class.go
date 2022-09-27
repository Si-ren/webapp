package class

import (
	"context"
	"github.com/sirupsen/logrus"
	classpb "grpcProtobuf/class/api/v1"
)

type Service struct {
}

func (s Service) GrabClass(c context.Context, req *classpb.GrabClassRequest) (*classpb.ClassResponse, error) {
	logrus.Info("Get GrabClassRequest successful")
	return &classpb.ClassResponse{IsGet: true}, nil
}
