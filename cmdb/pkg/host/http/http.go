package http

import (
	"cmdb/conf"
	"cmdb/pkg"
	"cmdb/pkg/host"
	"cmdb/pkg/host/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	api                = &handler{}
	_   pkg.GinService = (*handler)(nil)
)

type handler struct {
	svc host.Service
	log *logrus.Logger
}

func (h *handler) Name() string {
	return host.AppName
}

func NewHostHandler(svc host.Service) *handler {
	return &handler{
		svc: svc,
		log: conf.Log,
	}
}

func (h *handler) Configure() error {
	h.log = conf.Log
	if pkg.Host == nil {
		return fmt.Errorf("dependence service host not ready")
	}
	h.svc = pkg.Host
	return nil
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)

}

func init() {
	api = NewHostHandler(impl.HostService)
	pkg.GinRegister(api)
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
