package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type SlowQueryCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewSlowQueryCollector(db *sql.DB) *SlowQueryCollector {
	desc := prometheus.NewDesc("mysql_slow_query", "mysql slow query", nil, nil)
	return &SlowQueryCollector{
		baseCollector: newBaseCollector(db),
		desc:          desc,
	}
}

func (c *SlowQueryCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}
func (c *SlowQueryCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.status("slow_queries")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
