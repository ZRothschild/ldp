package user

import (
	"context"
	"github.com/ZRothschild/ldp/gen/register"
	"google.golang.org/grpc/grpclog"
)

type registerServer struct {
	register.UnsafeRegisterServiceServer
}

func NewRegisterServer() register.RegisterServiceServer {
	return new(registerServer)
}

func (s *registerServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *registerServer) Register(ctx context.Context, params *register.RegisterReq) (*register.RegisterResp, error) {
	var (
		resp = new(register.RegisterResp)
	)
	grpclog.Info(params, resp)
	return resp, nil
}
