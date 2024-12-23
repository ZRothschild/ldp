package user

import (
	"context"
	"github.com/ZRothschild/ldp/gen/user"
	"google.golang.org/grpc/grpclog"
)

type userServer struct {
	user.UnsafeUserServiceServer
}

func NewUserServer() user.UserServiceServer {
	return new(userServer)
}

func (s *userServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *userServer) UserDetail(ctx context.Context, params *user.UserReq) (*user.UserResp, error) {
	var (
		resp = new(user.UserResp)
	)
	grpclog.Info(params)
	return resp, nil
}
