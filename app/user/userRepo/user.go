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

type (
	FirstCondParams struct {
		Nickname string
		Password string
	}
)

func FirstCond(ctx context.Context, tx *gorm.DB, params FirstCondParams) *gorm.DB {
	tx.WithContext(ctx)
	if params.Nickname != "" {
		tx.Where("nickname = ?", params.Nickname)
	}
	if params.Password != "" {
		tx.Where("password = ?", params.Password)
	}
	return tx
}

func (r *UserRepo) Create(ctx context.Context, user *userM.User, tx *gorm.DB) error {
	if tx == nil {
		tx = r.DB.DB
	}
	return tx.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) First(ctx context.Context, params FirstCondParams, dest interface{}, tx *gorm.DB) error {
	if tx == nil {
		tx = r.DB.DB
	}
	return FirstCond(ctx, tx, params).First(dest).Error
}

func (r *UserRepo) Login(ctx context.Context, nickname, password string, dest interface{}) (err error) {
	if err = r.First(ctx, FirstCondParams{Nickname: nickname, Password: password}, dest, nil); err != nil {
		return err
	}
	return err
}
