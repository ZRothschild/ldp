package businessM

import "github.com/ZRothschild/ldp/app/base/baseM"

type (
	Business struct {
		baseM.BaseM
		BusinessName string `gorm:"type:varchar(255);column:business_name;not null;default:'';comment:业务名称" json:"businessName"`
	}
)
