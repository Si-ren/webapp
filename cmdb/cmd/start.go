package cmd

import (
	"cmdb/conf"
	"cmdb/pkg"
	"cmdb/pkg/host/impl"
	"cmdb/protocol"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "Demo后端API服务",
	Long:  `Demo后端API服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 初始化服务层 Ioc初始化
		if err := impl.Service.Config(); err != nil {
			return err
		}
		pkg.Host = impl.Service

		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svr, err := newService(conf.Configure())
		if err != nil {
			return err
		}

		// 等待信号处理
		go svr.waitSign(ch)
		// 启动服务
		if err := svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}
		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	svr := &service{
		http: http,
		log:  Log,
	}

	return svr, nil
}

var (
	Log      *logrus.Logger
	LogLevel map[string]logrus.Level = map[string]logrus.Level{
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
	}
)

type service struct {
	http *protocol.HTTPService
	log  *logrus.Logger
}

func (s *service) start() error {
	return s.http.Start()
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	case "etcd":
		return errors.New("not implemented")
	default:
		return errors.New("unknown config type")
	}
	return nil
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	lc := conf.Configure().Log
	Log.SetLevel(LogLevel[lc.Level])
	fmt.Println("log level: %s", lc.Level)

	switch lc.To {
	case conf.ToStdout:
		Log.Out = os.Stdout
	case conf.ToFile:
		file, err := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			Log.Out = file
		} else {
			Log.Info("Failed to log to file, using default stderr")
		}
	}
	switch lc.Format {
	case conf.JSONFormat:
		Log.Formatter = new(logrus.JSONFormatter)
	}
	Log.Info("Init log config complete!!")
	return nil
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:
			// 资源清理
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
			if err := s.http.Stop(); err != nil {
				s.log.Errorf("graceful shutdown err: %s, force exit", err)
			}
			s.log.Infof("service stop complete")
			return
		}
	}
}

func init() {
	RootCmd.AddCommand(serviceCmd)
}
