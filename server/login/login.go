package user

import (
	"context"
	"github.com/ZRothschild/ldp/gen/login"
	"google.golang.org/grpc/grpclog"
)

type loginServer struct {
	login.UnsafeLoginServiceServer
}

func NewLoginServer() login.LoginServiceServer {
	return new(loginServer)
}

func (s *loginServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *loginServer) Login(ctx context.Context, params *login.LoginReq) (*login.LoginResp, error) {
	var (
		resp = new(login.LoginResp)
	)
	grpclog.Info(params, resp)
	return resp, nil
}
