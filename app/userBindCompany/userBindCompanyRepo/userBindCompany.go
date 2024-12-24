package userBindCompanyRepo

import (
	"context"
	"github.com/ZRothschild/ldp/app/userBindCompany/userBindCompanyM"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	"gorm.io/gorm"
)

type UserBindCompanyRepo struct {
	*mysql.DB
}

func NewUserBindCompanyRepo(db *mysql.DB) *UserBindCompanyRepo {
	return &UserBindCompanyRepo{DB: db}
}

func (r *UserBindCompanyRepo) Create(ctx context.Context, userBindCompanyM *userBindCompanyM.UserBindCompany, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(userBindCompanyM).Error
	}
	return r.DB.WithContext(ctx).Create(userBindCompanyM).Error
}
