package impl

import (
	"cmdb/cmd"
	"cmdb/conf"
	"database/sql"
	"github.com/sirupsen/logrus"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db  *sql.DB
	log *logrus.Logger
}

func (s *service) Config() error {
	db, err := conf.Configure().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = cmd.Log
	s.db = db
	return nil
}
