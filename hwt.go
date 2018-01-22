package hwt

import (
	"context"
	fmt "fmt"

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

	r := &pb.HelloResp{
		Text:     fmt.Sprintf("%s said: %v", u, req.Subject),
		Hostname: s.Hostname,
	}
	return r, nil
}
