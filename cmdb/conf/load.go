package conf

import (
	"cmdb/cmd"
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

var (
	global   *Config
	Log      *logrus.Logger
	LogLevel map[string]logrus.Level = map[string]logrus.Level{
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
	}
)

// Configure 全局配置对象
func Configure() *Config {
	if global == nil {
		panic("Load Config first")
	}

	return global
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}

// config 为全局变量, 只需要load 即可全局可用户
func LoadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := LoadConfigFromToml(cmd.ConfFile)
		if err != nil {
			return err
		}
	case "env":
		err := LoadConfigFromEnv()
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
func LoadGlobalLogger() error {
	lc := Configure().Log
	Log.SetLevel(LogLevel[lc.Level])
	fmt.Println("log level: %s", lc.Level)

	switch lc.To {
	case ToStdout:
		Log.Out = os.Stdout
	case ToFile:
		file, err := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			Log.Out = file
		} else {
			Log.Info("Failed to log to file, using default stderr")
		}
	}
	switch lc.Format {
	case JSONFormat:
		Log.Formatter = new(logrus.JSONFormatter)
	}
	Log.Info("Init log config complete!!")
	return nil
}
