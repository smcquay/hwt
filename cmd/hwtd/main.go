package main

import (
	"log"
	"net/http"
	"os"

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
	if err := http.ListenAndServe(":8080", th); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
