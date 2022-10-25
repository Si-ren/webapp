package http

// import (
// 	"cmdb/cmd"
// 	"cmdb/pkg"
// 	"cmdb/pkg/host"
// 	"fmt"

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/sirupsen/logrus"
// )

// var (
// 	api = &handler{}
// )

// type handler struct {
// 	svc host.Service
// 	log *logrus.Logger
// }

// func (h *handler) Config() error {
// 	h.log = cmd.Log
// 	if pkg.Host == nil {
// 		return fmt.Errorf("dependence service host not ready")
// 	}
// 	h.service = pkg.Host
// 	return nil
// }

// func RegistAPI(r *httprouter.Router) {
// 	api.Config()
// 	r.GET("/hosts", api.QueryHost)
// 	r.POST("/hosts", api.CreateHost)
// 	r.GET("/hosts/:id", api.DescribeHost)
// 	r.DELETE("/hosts/:id", api.DeleteHost)
// 	r.PUT("/hosts/:id", api.PutHost)
// 	r.PATCH("/hosts/:id", api.PatchHost)
// }
