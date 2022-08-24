package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

func main1() {
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "cpu",
		Help:        "cpu total",
		ConstLabels: prometheus.Labels{"label": "test"},
	})
	//固定label
	disk := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "siri_disk_name",
		Help:        "siri disk help",
		ConstLabels: nil,
	}, []string{"mount"})
	disk.WithLabelValues("c:").Set(100)
	disk.WithLabelValues("d:").Set(200)

	request := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Request_Number",
		Help: "Http Get Count Number",
	})

	//设置指标的值
	cpu.Set(2)
	//在原有指标上加10
	cpu.Add(10)
	//注册指标
	prometheus.MustRegister(cpu)
	prometheus.MustRegister(disk)
	prometheus.MustRegister(request)
	request.Add(1)
	codeStatus := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_code_status",
		Help: "http code status",
	}, []string{"status"})
	prometheus.MustRegister(codeStatus)
	codeStatus.WithLabelValues("200").Add(10)
	codeStatus.WithLabelValues("300").Add(10)
	codeStatus.WithLabelValues("500").Add(10)

	go func() {
		for range time.Tick(time.Second) {
			disk.WithLabelValues("e:").Add(float64(rand.Int()))
		}
	}()

	//在metrics接口访问的时候,通过回调获取值
	call := prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "name",
		Help: "Help",
	}, func() float64 {
		fmt.Println("prometheus call")
		return rand.Float64()
	})
	prometheus.MustRegister(call)

	//暴露
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
