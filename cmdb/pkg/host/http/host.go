package http

import (
	"cmdb/pkg/host"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateHost(c *gin.Context) {
	ins := host.NewHost()
	// 解析参数

	if err := c.Bind(&ins); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"StatusInternalServerError": "FAILED"})
		return
	}

	c.JSON(http.StatusOK, ins)
}

func (h *handler) QueryHost(c *gin.Context) {
	//解析参数
	hostRequest := host.NewQueryHostRequestFromHTTP(c.Request)

	if hostSet, err := h.svc.QueryHost(c.Request.Context(), hostRequest); err != nil {
		h.log.Errorf("QueryHost err: %s ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"StatusBadRequest": err.Error()})

	} else {
		c.JSON(http.StatusOK, hostSet)
	}
}
