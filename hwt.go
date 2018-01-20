package hwt

import (
	"context"
	fmt "fmt"

	"github.com/twitchtv/twirp"
	pb "mcquay.me/hwt/rpc/hwt"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	if req.Subject == "" {
		return nil, twirp.RequiredArgumentError("subject")
	}
	return &pb.HelloResp{Text: fmt.Sprintf("echo: %v", req.Subject)}, nil
}
