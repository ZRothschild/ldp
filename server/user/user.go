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

func (s *userServer) Echo(ctx context.Context, msg *user.StringMessage) (*user.StringMessage, error) {
	grpclog.Info(msg)
	return msg, nil
}
