package grpc_resource

import (
	"cmdb/pkg"
	"database/sql"

	"google.golang.org/grpc"
)

var (
	svc = &service{}
	// 静态检查
	_ pkg.GrpcService = (*service)(nil)
)

type service struct {
	db *sql.DB
}

func (s *service) Registry(g *grpc.Server) {

}

func (s *service) Config() error {
	return nil
}

func (s *service) Name() string {
	return " "
}

func init() {

}
