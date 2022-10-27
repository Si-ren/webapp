package http

import (
	"cmdb/pkg/host"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateHost(c *gin.Context) {
	ins := host.NewHost()
	fmt.Println(c.body)
	// 解析参数
	if err := c.Bind(&ins); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"ERROR": " CreateHost failed"})
		return
	}

	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"StatusInternalServerError": "FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"create host": "OK"})
}
