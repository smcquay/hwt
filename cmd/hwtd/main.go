package main

import (
	"log"
	"net/http"

	"mcquay.me/hwt"
)

func main() {
	s := &hwt.Server{}
	th := hwt.NewHelloWorldServer(s, nil)
	if err := http.ListenAndServe(":8080", th); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
