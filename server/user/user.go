package user

import (
	"context"
	"github.com/ZRothschild/ldp/app/user/userRepo"
	"github.com/ZRothschild/ldp/gen/user"
	"google.golang.org/grpc/grpclog"
)

type userServer struct {
	user.UnsafeUserServiceServer
	*userRepo.UserRepo
}

func NewUserServer(userR *userRepo.UserRepo) user.UserServiceServer {
	return &userServer{
		UserRepo: userR,
	}
}

func (s *userServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *userServer) UserDetail(ctx context.Context, params *user.UserReq) (*user.UserResp, error) {
	var (
		resp = new(user.UserResp)

	)

	s.UserRepo.ById(ctx，)

	grpclog.Info(params)
	return resp, nil
}
