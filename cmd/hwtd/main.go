package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"mcquay.me/hwt"
	pb "mcquay.me/hwt/rpc/hwt"
)

func main() {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalf("cannot get hostname: %v", err)
	}
	s := &hwt.Server{hn}
	th := pb.NewHelloWorldServer(s, nil)
	sm := http.NewServeMux()
	sm.Handle("/", th)
	sm.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", sm); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
