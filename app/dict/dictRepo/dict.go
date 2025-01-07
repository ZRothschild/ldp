package dictRepo

import (
	"context"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	"github.com/ZRothschild/ldp/infrastr/static/entry"
	"gorm.io/gorm"
	"math"
)

type DictRepo struct {
	*mysql.DB
}

func NewDictRepo(db *mysql.DB) *DictRepo {
	return &DictRepo{DB: db}
}

type (
	FirstParams struct {
		Nickname string
		Password string
	}
)

func FirstCond(ctx context.Context, tx *gorm.DB, params FirstParams) *gorm.DB {
	tx.WithContext(ctx)
	if params.Nickname != "" {
		tx.Where("nickname = ?", params.Nickname)
	}
	if params.Password != "" {
		tx.Where("password = ?", params.Password)
	}
	return tx
}

func (r *DictRepo) First(ctx context.Context, params FirstParams, dest interface{}, tx *gorm.DB) error {
	if tx == nil {
		tx = r.DB.DB
	}
	return FirstCond(ctx, tx, params).First(dest).Error
}

type (
	FindParams struct {
		Nickname string
		Password string
		entry.Pagination
	}
)

func FidCond(ctx context.Context, tx *gorm.DB, params *FindParams) *gorm.DB {
	tx.WithContext(ctx)
	if params.Nickname != "" {
		tx.Where("nickname = ?", params.Nickname)
	}
	if params.Password != "" {
		tx.Where("password = ?", params.Password)
	}
	return tx
}

func (r *DictRepo) Find(ctx context.Context, params *FindParams, dest interface{}, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = r.DB.DB
	}
	tx = FidCond(ctx, tx, params)
	if params.Pagination.PageSize > 0 {
		if err = tx.Count(&params.Pagination.Total).Error; err != nil {
			return err
		}
		totalPage := int(math.Ceil(float64(params.Pagination.Total) / float64(params.Pagination.PageSize)))
		if totalPage == 0 {
			return err
		} else if params.Pagination.Page > totalPage {
			params.Pagination.Page = totalPage
		}
		tx.Offset((params.Pagination.Page - 1) * params.Pagination.PageSize).Limit(params.Pagination.PageSize)
	}
	return tx.Find(dest).Error
}
