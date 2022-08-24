package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"promAgent/config"
)

func init() {
	//从环境变量读取,优先级比文件高
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PROM_AGENT")
}

//viper配置文件读取
func InitConfig(path string) *config.AgentConfig {
	var config config.AgentConfig
	//从配置文件读取
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
	fmt.Println(config.Addr)

	//读取uuid,cmdb用来分辨node
	uuidPath := "promAgent.uuid"
	if ctx, err := ioutil.ReadFile(uuidPath); err == nil {
		config.UUID = string(ctx)
	} else if os.IsNotExist(err) {
		config.UUID = uuid.NewString()
		ioutil.WriteFile(uuidPath, []byte(config.UUID), os.ModePerm)
	} else {
		logrus.Fatal(err)
	}
	return &config
}
