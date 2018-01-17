package hwt

import (
	"context"
	fmt "fmt"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	return &HelloResp{Text: fmt.Sprintf("echo: %v", req.Subject)}, nil
}
