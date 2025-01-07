package dictM

import "github.com/ZRothschild/ldp/app/base/baseM"

type (
	Dict struct {
		baseM.BaseM
		Pid    uint64 `gorm:"type:bigint UNSIGNED;column:pid;not null;default:0;comment:pid" json:"pid"`
		Mark   string `gorm:"type:varchar(255);column:mark;not null;default:'';comment:标识" json:"mark"`
		Name   string `gorm:"type:varchar(255);column:name;not null;default:'';comment:名称" json:"name"`
		EnName string `gorm:"type:varchar(255);column:en_name;not null;default:'';comment:英文名称" json:"enName"`
	}
)
