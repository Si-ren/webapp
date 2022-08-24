package config

import "time"

type AgentConfig struct {
	UUID        string        `mapstructure:"uuid"`
	Addr        string        `mapstructure:"addr"`
	SeverConfig *ServerConfig `mapstructure:"server"`
	LogConfig   *LogConfig    `mapstructure:"log"`
	TaskConfig  *TaskConfig   `mapstructure:"task"`
}

type ServerConfig struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

type LogConfig struct {
	FileName   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	Compress   bool   `mapstructure:"compress"`
}

type TaskConfig struct {
	Register *RegisterConfig `mapstructure:"register"`
}

type RegisterConfig struct {
	Interval time.Duration `mapstructure:"interval"`
}
