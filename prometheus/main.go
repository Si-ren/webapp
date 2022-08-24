package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"prometheus/collectors"
)

func main() {
	addr := ":9999"
	dsn := "root:root@tcp(localhost:3306)/cmdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	//定义指标

	//注册指标
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewSlowQueryCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	prometheus.MustRegister(collectors.NewConnectCollector(db))
	//注册控制器
	http.Handle("/metrics", promhttp.Handler())
	//启动web服务
	http.ListenAndServe(addr, nil)
}
