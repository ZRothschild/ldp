package user

import (
	"context"
	"github.com/ZRothschild/ldp/app/base/baseM"
	"github.com/ZRothschild/ldp/app/company/companyM"
	"github.com/ZRothschild/ldp/app/company/companyRepo"
	"github.com/ZRothschild/ldp/app/user/userM"
	"github.com/ZRothschild/ldp/app/user/userRepo"
	"github.com/ZRothschild/ldp/app/userBindCompany/userBindCompanyM"
	"github.com/ZRothschild/ldp/app/userBindCompany/userBindCompanyRepo"
	"github.com/ZRothschild/ldp/gen/register"
	"google.golang.org/grpc/grpclog"
)

type registerServer struct {
	register.UnsafeRegisterServiceServer
	*companyRepo.CompanyRepo
	*userRepo.UserRepo
	*userBindCompanyRepo.UserBindCompanyRepo
}

func NewRegisterServer(userR *userRepo.UserRepo, companyR *companyRepo.CompanyRepo, userBindCompanyR *userBindCompanyRepo.UserBindCompanyRepo) register.RegisterServiceServer {
	return &registerServer{
		UnsafeRegisterServiceServer: register.UnimplementedRegisterServiceServer{},
		UserRepo:                    userR,
		CompanyRepo:                 companyR,
		UserBindCompanyRepo:         userBindCompanyR,
	}
}

func (s *registerServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

// Register 用户登陆
func (s *registerServer) Register(ctx context.Context, params *register.RegisterReq) (*register.RegisterResp, error) {
	var (
		err      error
		tx       = s.UserRepo.DB.WithContext(ctx).Begin()
		resp     = new(register.RegisterResp)
		userInfo = &userM.User{
			Nickname: params.GetNickname(),
			Username: params.GetUsername(),
			Password: params.GetPassword(),
		}
		companyInfo = new(companyM.Company)
	)

	if params.GetRegisterType() == register.RegisterType_User {
		userInfo = &userM.User{
			Nickname:    params.GetNickname(),
			Username:    params.GetUsername(),
			Password:    params.GetPassword(),
			Email:       params.GetEmail(),
			CompanyName: params.GetCompanyName(),
			Phone:       params.GetPhone(),
			Mobile:      params.GetMobile(),
			IdCardFront: params.GetIdCardFront(),
			IdCardBack:  params.GetIdCardBack(),
			Avatar:      params.GetAvatar(),
			Seniority:   params.GetSeniority(),
			Profile:     params.GetProfile(),
			Location:    params.GetLocation(),
		}
	}

	if err = s.UserRepo.Create(ctx, userInfo, tx); err != nil {
		return resp, err
	}

	base := baseM.BaseM{
		CreatedBy: userInfo.ID,
		UpdatedBy: userInfo.ID,
	}
	
	if params.GetRegisterType() == register.RegisterType_Company {
		companyInfo = &companyM.Company{
			BaseM:       base,
			CompanyName: params.GetCompanyName(),
			Phone:       params.GetPhone(),
			Mobile:      params.GetMobile(),
			IdCardFront: params.GetIdCardFront(),
			IdCardBack:  params.GetIdCardBack(),
			Avatar:      params.GetAvatar(),
			Seniority:   params.GetSeniority(),
			Profile:     params.GetProfile(),
			Location:    params.GetLocation(),
			License:     params.GetLicense(),
		}
		if err = s.CompanyRepo.Create(ctx, companyInfo, tx); err != nil {
			return resp, err
		}
		UserBindCompanyInfo := &userBindCompanyM.UserBindCompany{
			BaseM:        base,
			UserId:       userInfo.ID,
			CompanyId:    companyInfo.ID,
			Relationship: userBindCompanyM.SuperRelationship,
		}
		if err = s.UserBindCompanyRepo.Create(ctx, UserBindCompanyInfo, tx); err != nil {
			return resp, err
		}
	}

	grpclog.Info(params, resp)
	return resp, err
}
