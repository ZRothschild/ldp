package companyMd

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	Company struct {
		baseM.BaseM
		CompanyName string `gorm:"type:varchar(255);column:company_name;not null;default:'';comment:公司" json:"companyName"`
		Phone       string `gorm:"type:varchar(255);column:phone;not null;default:'';comment:座机号" json:"phone"`
		Mobile      string `gorm:"type:varchar(255);column:mobile;not null;default:'';comment:手机号" json:"mobile"`
		IdCardFront string `gorm:"type:varchar(255);column:id_card_front;not null;default:'';comment:法人身份证正面" json:"idCardFront"`
		IdCardBack  string `gorm:"type:varchar(255);column:id_card_back;not null;default:'';comment:法人身份证反面" json:"idCardBack"`
		Avatar      string `gorm:"type:varchar(255);column:avatar;not null;default:'';comment:公司logo" json:"avatar"`
		Seniority   int    `gorm:"type:int;column:seniority;not null;default:0;comment:公司年龄" json:"seniority"`
		Profile     string `gorm:"type:varchar(255);column:profile;not null;default:'';comment:公司简介" json:"profile"`
		Location    string `gorm:"type:varchar(255);column:location;not null;default:'';comment:公司地址" json:"location"`
		License     string `gorm:"type:varchar(255);column:license;not null;default:'';comment:营业执照" json:"license"`
	}
)
