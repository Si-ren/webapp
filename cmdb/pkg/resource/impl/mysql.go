package impl

import (
	"cmdb/pkg/host"
	"database/sql"

	"github.com/sirupsen/logrus"
)

var ()

type service struct {
	db   *sql.DB
	log  logrus.Logger
	host host.Service
}
