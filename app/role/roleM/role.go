package roleM

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	Role struct {
		baseM.BaseM
		RoleName string `gorm:"type:varchar(255);column:role_name;not null;default:'';comment:角色名称" json:"roleName"`
	}
)
