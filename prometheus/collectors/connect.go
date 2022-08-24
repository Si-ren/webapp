package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type ConnectCollector struct {
	*baseCollector
	maxConnectedDesc    *prometheus.Desc
	threadConnectedDesc *prometheus.Desc
}

func NewConnectCollector(db *sql.DB) *ConnectCollector {
	maxConnected := prometheus.NewDesc("maxConnected", "Max Connected Count", nil, nil)
	threadConnected := prometheus.NewDesc("threadConnected", "Thread Connected Count", nil, nil)
	return &ConnectCollector{
		baseCollector:       newBaseCollector(db),
		maxConnectedDesc:    maxConnected,
		threadConnectedDesc: threadConnected,
	}
}

func (c *ConnectCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.maxConnectedDesc
	desc <- c.threadConnectedDesc
}

func (c *ConnectCollector) Collect(metrics chan<- prometheus.Metric) {
	metrics <- prometheus.MustNewConstMetric(c.maxConnectedDesc, prometheus.CounterValue, c.variables("max_connections"))
	metrics <- prometheus.MustNewConstMetric(c.threadConnectedDesc, prometheus.CounterValue, c.status("Threads_connected"))
}
