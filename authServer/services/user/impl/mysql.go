package impl

import (
	"cmdb/conf"
	"cmdb/services"
	"cmdb/services/user"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	queryUserSQL      = `SELECT * FROM resource as r LEFT JOIN user h ON r.id=h.resource_id`
	queryResourceSQL  = `SELECT * FROM resource where resource_id >= ? limit ?;`
	queryDescribeSQL  = "SELECT * FROM `describe` where  describe_id >= ? limit ?;"
	queryBaseSQL      = `SELECT base_id,instance_id,sync_at,vendor,region,zone,create_at,resource_hash,describe_hash FROM base where base_id >= ? limit ?;`
	deleteUserSQL     = `DELETE FROM user WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

var (
	// UserService 服务实例
	UserService              = &service{}
	_           user.Service = (*service)(nil)
)

type service struct {
	db  *gorm.DB
	log *logrus.Logger
}

func (s *service) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) QueryUser(ctx context.Context, request *user.QueryUserRequest) (*user.UserSet, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DescribeUser(ctx context.Context, request *user.DescribeUserRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Config() error {
	db, err := conf.Configure().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = conf.Log
	s.db = db
	return nil
}

func init() {
	services.ImplRegister(UserService)
}
