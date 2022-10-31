package conf

import "github.com/sirupsen/logrus"

// LogFormat 日志格式
type LogFormat string

var (
	Log *logrus.Logger = logrus.New()
)

const (
	// TextFormat 文本格式
	TextFormat = LogFormat("text")
	// JSONFormat json格式
	JSONFormat = LogFormat("json")
)

// LogTo 日志记录到哪儿
type LogTo string

const (
	// ToFile 保存到文件
	ToFile = LogTo("file")
	// ToStdout 打印到标准输出
	ToStdout = LogTo("stdout")
)
