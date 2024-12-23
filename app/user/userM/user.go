package userM

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	User struct {
		baseM.BaseM
		Mobile       string `gorm:"type:varchar(255);column:mobile;not null;default:'';comment:用户手机号" json:"mobile"`
		Email        string `gorm:"type:varchar(255);column:email;not null;default:'';comment:用户邮箱" json:"email"`
		Username     string `gorm:"type:varchar(255);column:username;not null;default:'';comment:用户真实名称" json:"username"`
		Nickname     string `gorm:"type:varchar(255);column:nickname;not null;default:'';comment:用户昵称" json:"nickname"`
		Password     string `gorm:"type:varchar(255);column:password;not null;default:'';comment:用户密码" json:"password"`
		IdCardFront  string `gorm:"type:varchar(255);column:id_card_front;not null;default:'';comment:身份证正面" json:"idCardFront"`
		IdCardBack   string `gorm:"type:varchar(255);column:id_card_back;not null;default:'';comment:身份证反面" json:"idCardBack"`
		Avatar       string `gorm:"type:varchar(255);column:avatar;not null;default:'';comment:用户头像" json:"avatar"`
		Seniority    int    `gorm:"type:int;column:seniority;not null;default:0;comment:用户工龄" json:"seniority"`
		Profile      string `gorm:"type:varchar(255);column:profile;not null;default:'';comment:用户简介" json:"profile"`
		WorkLocation string `gorm:"type:varchar(255);column:work_location;not null;default:'';comment:工作地址" json:"workLocation"`
		WorkCompany  string `gorm:"type:varchar(255);column:work_company;not null;default:'';comment:工作单位" json:"workCompany"`
	}
)
