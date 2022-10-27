package http

import (
	"cmdb/pkg/host"
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

func NewHostHandler(svc host.Service, log *logrus.Logger) *handler {
	return &handler{
		svc: svc,
		log: log,
	}
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
