package protocol

import (
	"cmdb/conf"
	"cmdb/services"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

// HTTPService http服务
type HTTPService struct {
	//r      *httprouter.Router
	r gin.IRouter
	l *logrus.Logger

	c      *conf.Config
	server *http.Server
}

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	//r := httprouter.New()
	r := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.Configure().App.Addr(),
		Handler:           r,
	}
	return &HTTPService{
		r:      r,
		server: server,
		l:      conf.Log,
		c:      conf.Configure(),
	}
}

// Start 启动服务
func (s *HTTPService) Start() error {
	//初始化service实现
	services.InitConfigSvc()
	// 初始化service路由
	services.InitGinSvc(s.r)

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
