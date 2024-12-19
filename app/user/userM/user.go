package userM

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	User struct {
		baseM.BaseM
		Email    string `gorm:"type:varchar(255);column:email;not null;default:'';comment:用户邮箱" json:"email"`
		Username string `gorm:"type:varchar(255);column:username;not null;default:'';comment:用户真实名称" json:"username"`
		Nickname string `gorm:"type:varchar(255);column:nickname;not null;default:'';comment:用户昵称" json:"nickname"`
	}
)
