package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	pb "mcquay.me/hwt/rpc/hwt"
)

const usage = "hwtc <server address> [subject]"

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%v\n", usage)
		os.Exit(1)
	}

	c := pb.NewHelloWorldProtobufClient(fmt.Sprintf("http://%s", os.Args[1]), &http.Client{})

	resp, err := c.Hello(context.Background(), &pb.HelloReq{Subject: strings.Join(os.Args[2:], " ")})
	if err != nil {
		fmt.Fprintf(os.Stderr, "hello: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", resp)
}
