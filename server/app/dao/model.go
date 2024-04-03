package dao

import (
	"time"

	"gorm.io/gorm"
)

// 通用结构体
type Model struct {
	gorm.Model
	Id        uint           `gorm:"column:id;type:int;primary_key;unsigned;not null;autoIncrement" json:"id"` // Id
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp" json:"createdAt"`                        // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`                        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`                        // 删除时间
}

var (
	// 客户端对象
	DB *gorm.DB
	// 初始化错误
	err error
	// 自动创表列表
	dst = make([]interface{}, 0)
)
