package company

import (
	"context"
	"github.com/ZRothschild/ldp/gen/company"
	"google.golang.org/grpc/grpclog"
)

type companyServer struct {
	company.UnsafeCompanyServiceServer
}

func NewCompanyServer() company.CompanyServiceServer {
	return new(companyServer)
}

func (s *companyServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *companyServer) CompanyDetail(ctx context.Context, params *company.CompanyReq) (*company.CompanyResp, error) {
	var (
		resp = new(company.CompanyResp)
	)
	grpclog.Info(params)
	return resp, nil
}
