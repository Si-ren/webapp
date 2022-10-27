package http

import (
	"cmdb/cmd"
	"cmdb/pkg"
	"cmdb/pkg/host"
	"cmdb/pkg/host/impl"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	HostHandler = &handler{}
)

type handler struct {
	svc host.Service
	log *logrus.Logger
}

func NewHostHandler(svc host.Service) *handler {
	return &handler{
		svc: svc,
		log: cmd.Log,
	}
}

func (h *handler) Config() error {
	h.log = cmd.Log
	if pkg.Host == nil {
		return fmt.Errorf("dependence service host not ready")
	}
	h.svc = impl.Service
	return nil
}

func (h *handler) RegistryApi(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
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
