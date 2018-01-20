package hwt

import (
	"context"
	fmt "fmt"

	pb "mcquay.me/hwt/rpc/hwt"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: fmt.Sprintf("echo: %v", req.Subject)}, nil
}
