package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type TrafficCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewTrafficCollector(db *sql.DB) *TrafficCollector {
	desc := prometheus.NewDesc("mysql_traffic", "mysql traffic", []string{"direction"}, nil)
	return &TrafficCollector{
		baseCollector: newBaseCollector(db),
		desc:          desc,
	}
}

func (c *TrafficCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}
func (c *TrafficCollector) Collect(metrics chan<- prometheus.Metric) {

	in := c.status("Bytes_received")

	out := c.status("Bytes_sent")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, in, "in")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, out, "out")
}
