package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/twitchtv/twirp"

	"mcquay.me/hwt"
	pb "mcquay.me/hwt/rpc/hwt"
)

const usage = "hwtc <server address> [subject]"

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%v\n", usage)
		os.Exit(1)
	}

	c := pb.NewHelloWorldProtobufClient(fmt.Sprintf("http://%s", os.Args[1]), &http.Client{})

	h := http.Header{}
	h.Set("sm-auth", hwt.PSK)
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, h)
	if err != nil {
		fmt.Fprintf(os.Stderr, "setting twirp headers: %v\n", err)
		os.Exit(1)
	}

	for i := 0; ; i++ {
		resp, err := c.Hello(ctx, &pb.HelloReq{Subject: strings.Join(os.Args[2:], " ")})
		if err != nil {
			fmt.Fprintf(os.Stderr, "hello: %#v\n", err)
			os.Exit(1)
		}
		fmt.Printf("0x%08x: %#v\n", i, resp)
	}
}
