package companyRepo

import (
	"context"
	"github.com/ZRothschild/ldp/app/company/companyM"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	"gorm.io/gorm"
)

type CompanyRepo struct {
	*mysql.DB
}

func NewCompanyRepo(db *mysql.DB) *CompanyRepo {
	return &CompanyRepo{DB: db}
}

func (r *CompanyRepo) Create(ctx context.Context, company *companyM.Company, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(company).Error
	}
	return r.DB.Create(company).Error
}
