package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"promAgent/config"
)

func InitLog(verbose bool, config *config.AgentConfig) {
	logger := &lumberjack.Logger{
		Filename:   config.LogConfig.FileName,
		MaxSize:    config.LogConfig.MaxSize,
		MaxBackups: config.LogConfig.MaxBackups,
		Compress:   config.LogConfig.Compress,
	}
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		//打印方法
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(logger)
	}
}
