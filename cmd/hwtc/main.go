package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"mcquay.me/hwt"
)

const usage = "hwtc [subject]"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%v\n", usage)
		os.Exit(1)
	}

	c := hwt.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := c.Hello(context.Background(), &hwt.HelloReq{Subject: strings.Join(os.Args[1:], " ")})
	if err != nil {
		fmt.Fprintf(os.Stderr, "hello: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", resp)
}
