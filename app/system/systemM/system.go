package systemM

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	System struct {
		baseM.BaseM
		SystemName string `gorm:"type:varchar(255);column:system_name;not null;default:'';comment:系统名称" json:"systemName"`
	}
)
