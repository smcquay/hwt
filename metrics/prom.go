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

			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.95: 0.001,
				0.99: 0.001,
				1.0:  0.0001,
			},
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
