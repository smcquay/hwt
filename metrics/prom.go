package metrics

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpReqLat = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "hwt_request_latency_ms",
			Help: "Latency in ms of http requests grouped by req path",
		},
		[]string{"path"},
	)
)

func RegisterPromMetrics() error {
	if err := prometheus.Register(httpReqLat); err != nil {
		return errors.Wrap(err, "registering http request latency")
	}
	return nil
}
