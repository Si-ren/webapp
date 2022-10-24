package plugins

import (
	"fmt"
	"github.com/imroc/req/v3"
	"os"
	"promAgent/config"
	"strings"
	"time"
)

type Register struct {
	config *config.AgentConfig
}

func (r *Register) Init(config *config.AgentConfig) {
	r.config = config
}

type NodeRegisterForm struct {
	UUID     string `json:"uuid"`
	Hostname string `json:"hostname"`
	Addr     string `json:"addr"`
}

type Target struct {
	Addr string `json:"addr"`
}

type Job struct {
	Name    string    `json:"name"`
	Targets []*Target `json:"targets"`
}

func (r *Register) Run() {
	ticker := time.NewTicker(r.config.TaskConfig.Register.Interval)
	defer ticker.Stop()
	//UUID
	//Hostname
	hostname, _ := os.Hostname()
	//Addr

	client := req.C().DevMode()

	form := &NodeRegisterForm{
		UUID:     r.config.UUID,
		Hostname: hostname,
		Addr:     r.config.Addr,
	}

	job := &Job{}

	for {
		//先执行
		fmt.Println("Send register info to cmdb-test")
		resp, err := client.R().SetHeader("Authorization", fmt.Sprintf("Token %s", r.config.SeverConfig.Token)).
			//SetQueryParam("uuid", r.config.UUID).
			//SetQueryParam("addr", r.config.Addr).
			//SetQueryParam("hostname", hostname).
			//SetBodyJsonString(`{"uuid": r.config.UUID}`).
			SetBody(form).
			Post(fmt.Sprintf("%s/v1/prometheus/register", strings.TrimRight(r.config.SeverConfig.Addr, "/")))
		if err != nil {
			fmt.Println("err:", err)
			if resp.Dump() != "" {
				fmt.Println("raw content:")
				fmt.Println(resp.Dump())
			}
			return
		}
		if !resp.IsSuccess() { // Status code not beetween 200 and 299
			fmt.Println("bad status:", resp.Status)
			fmt.Println("raw content:")
			fmt.Println(resp.Dump())
			return
		} else if resp.IsSuccess() {
			fmt.Println("Register success ...", job)
			resp.Unmarshal(job)
			fmt.Println(job)
		}

		<-ticker.C
		//return
	}

}
