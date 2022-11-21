package http

import (
	"cmdb/conf"
	"cmdb/services"
	"cmdb/services/user"
	"cmdb/services/user/impl"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	api                     = &handler{}
	_   services.GinService = (*handler)(nil)
)

type handler struct {
	svc user.Service
	log *logrus.Logger
}

func (h *handler) Name() string {
	return user.AppName
}

func NewHostHandler(svc user.Service) *handler {
	return &handler{
		svc: svc,
		log: conf.Log,
	}
}

func (h *handler) Configure() error {
	h.log = conf.Log
	if services.Host == nil {
		return fmt.Errorf("dependence service user not ready")
	}
	h.svc = services.Host
	return nil
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
	r.GET("/hosts", h.QueryHost)

}

func init() {
	api = NewHostHandler(impl.HostService)
	services.GinRegister(api)
}

// func RegistAPI(r *httprouter.Router) {
// 	HostHandler.Config()
// 	r.GET("/hosts", HostHandler.svc.CreateHost)
// 	r.POST("/hosts", HostHandler.CreateHost)
// 	r.GET("/hosts/:id", HostHandler.DescribeHost)
// 	r.DELETE("/hosts/:id", HostHandler.DeleteHost)
// 	r.PUT("/hosts/:id", HostHandler.PutHost)
// 	r.PATCH("/hosts/:id", HostHandler.PatchHost)
// }
