package user

import (
	"context"
	"github.com/ZRothschild/ldp/app/user/userRepo"
	"github.com/ZRothschild/ldp/gen/login"
	"google.golang.org/grpc/grpclog"
)

type loginServer struct {
	login.UnsafeLoginServiceServer
	*userRepo.UserRepo
}

func NewLoginServer(userR *userRepo.UserRepo) login.LoginServiceServer {
	return &loginServer{
		UserRepo: userR,
	}
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
