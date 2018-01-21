package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"mcquay.me/hwt"
	"mcquay.me/hwt/metrics"
	pb "mcquay.me/hwt/rpc/hwt"
)

func main() {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalf("cannot get hostname: %v", err)
	}

	if err := metrics.RegisterPromMetrics(); err != nil {
		log.Fatalf("registering prom metrics: %v", err)
	}

	s := &hwt.Server{hn}
	hs := hwt.NewMetricsHooks(metrics.HTTPLatency)
	th := pb.NewHelloWorldServer(s, hs)
	sm := http.NewServeMux()
	sm.Handle("/", th)
	sm.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", sm); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
