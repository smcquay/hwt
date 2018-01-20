package main

import (
	"log"
	"net/http"

	"mcquay.me/hwt"
	pb "mcquay.me/hwt/rpc/hwt"
)

func main() {
	s := &hwt.Server{}
	th := pb.NewHelloWorldServer(s, nil)
	if err := http.ListenAndServe(":8080", th); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
