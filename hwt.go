package hwt

import (
	"context"
	fmt "fmt"
	"go/src/math/rand"
	"time"

	"github.com/twitchtv/twirp"
	pb "mcquay.me/hwt/rpc/hwt"
)

type Server struct {
	Hostname string
}

func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	if req.Subject == "" {
		return nil, twirp.RequiredArgumentError("subject")
	}

	u, err := getUser(ctx)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if rand.Int()%100 == 0 {
		rest := time.Duration(rand.Intn(50)) * time.Millisecond
		time.Sleep(rest)
	}

	switch rand.Int() % 100 {
	case 0:
		return nil, twirp.NewError(twirp.Internal, "bleeding")
	case 1:
		return nil, twirp.NewError(twirp.ResourceExhausted, "some exhaustion")
	}

	r := &pb.HelloResp{
		Text:     fmt.Sprintf("%s said: %v", u, req.Subject),
		Hostname: s.Hostname,
	}
	return r, nil
}
