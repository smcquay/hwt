package metrics

import (
	"time"
)

func HTTPLatency(path string, dur time.Duration) {
	httpReqLat.WithLabelValues(path).Observe(float64(dur) / float64(time.Millisecond))
}

func HTTPCode(path string, status string) {
	httpStatus.WithLabelValues(path, status).Inc()
}
