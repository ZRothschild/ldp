package UserBindCompanyM

import "github.com/ZRothschild/ldp/app/base/baseM"

type (
	UserBindCompany struct {
		baseM.BaseM
		UserId       uint64 `gorm:"type:bigint UNSIGNED not null;column:user_id;comment:用户id" json:"userId"`
		CompanyId    uint64 `gorm:"type:bigint UNSIGNED not null;column:company_id;comment:公司id" json:"companyId"`
		Relationship int    `gorm:"type:int;column:relationship;not null;default:0;comment:用户与公司关系 100公司超级管理员 200 公司普通员工" json:"relationship"`
	}
)
