package baseM

import (
	"gorm.io/gorm"
	"time"
)

type (
	BaseM struct {
		ID        uint64         `gorm:"type:bigint UNSIGNED not null;column:id;autoIncrement;primarykey" json:"id"`
		CreatedAt time.Time      `gorm:"type:datetime;column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
		UpdatedAt time.Time      `gorm:"type:datetime;column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
		CreatedBy uint64         `gorm:"type:bigint UNSIGNED not null;column:created_by;default:0;comment:创建人" json:"createdBy"`
		UpdatedBy uint64         `gorm:"type:bigint UNSIGNED not null;column:updated_by;default:0;comment:更新人" json:"updatedBy"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	}
)
