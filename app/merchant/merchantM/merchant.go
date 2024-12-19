package merchantM

import (
	"github.com/ZRothschild/ldp/app/base/baseM"
)

type (
	Merchant struct {
		baseM.BaseM
		MerchantName string `gorm:"type:varchar(255);column:merchant_name;not null;default:'';comment:商户名称" json:"merchantName"`
	}
)
