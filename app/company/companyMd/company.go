package companyMd

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	Company struct {
		baseM.BaseM
		CompanyName     string `gorm:"type:varchar(255);column:company_name;not null;default:'';comment:公司" json:"companyName"`
		IdCardFront     string `gorm:"type:varchar(255);column:id_card_front;not null;default:'';comment:法人身份证正面" json:"idCardFront"`
		IdCardBack      string `gorm:"type:varchar(255);column:id_card_back;not null;default:'';comment:法人身份证反面" json:"idCardBack"`
		Avatar          string `gorm:"type:varchar(255);column:avatar;not null;default:'';comment:公司logo" json:"avatar"`
		Seniority       int    `gorm:"type:int;column:seniority;not null;default:0;comment:公司年龄" json:"seniority"`
		Profile         string `gorm:"type:varchar(255);column:profile;not null;default:'';comment:公司简介" json:"profile"`
		CompanyLocation string `gorm:"type:varchar(255);column:work_company;not null;default:'';comment:公司地址" json:"workCompany"`
		BusinessLicense string `gorm:"type:varchar(255);column:business_license;not null;default:'';comment:营业执照" json:"businessLicense"`
	}
)
