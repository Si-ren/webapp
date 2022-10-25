package http

import (
	"cmdb/cmd"
	"cmdb/pkg"
	"cmdb/pkg/host"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	api = &handler{}
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
	h.svc = pkg.Host
	return nil
}

func (h *handler) RegistryApi(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
}
