package user

import (
	"context"
	"crypto/md5"
	"github.com/ZRothschild/ldp/app/user/userM"
	"github.com/ZRothschild/ldp/app/user/userRepo"
	"github.com/ZRothschild/ldp/gen/login"
	"github.com/ZRothschild/ldp/infrastr/conf"
	"github.com/ZRothschild/ldp/infrastr/lib/jwt"
	"github.com/ZRothschild/ldp/infrastr/static/code"
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
		err      error
		userInfo = new(userM.User)
		resp     = &login.LoginResp{
			Success: true,
			Data:    new(login.LoginDetail),
		}
		md5pwd = md5.New().Sum([]byte(params.GetPassword()))
	)

	// 查询用户信息是否在数据库
	if err = s.UserRepo.Login(ctx, params.GetNickname(), string(md5pwd), userInfo); err != nil {
		return resp, code.PwdOrNicknameMatchErr
	}

	if resp.Data.Token, err = jwt.SignJwt(conf.Conf.JwtSk, jwt.UserInfo{Id: userInfo.ID, Nickname: userInfo.Nickname}); err != nil {
		return resp, err
	}
	grpclog.Info(params, resp)
	return resp, nil
}
