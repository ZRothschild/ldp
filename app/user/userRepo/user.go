package userRepo

import (
	"context"
	"github.com/ZRothschild/ldp/app/user/userM"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	"gorm.io/gorm"
)

type UserRepo struct {
	*mysql.DB
}

func NewUserRepo(db *mysql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *userM.User, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(user).Error
	}
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) First(ctx context.Context, user *userM.User, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(user).Error
	}
	return r.DB.WithContext(ctx).Create(user).Error
}
