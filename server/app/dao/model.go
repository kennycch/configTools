package dao

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 通用结构体
type Model struct {
	Id        uint32         `gorm:"autoIncrement" json:"id"` // Id
	CreatedAt time.Time      `json:"createdAt"`               // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`               // 更新时间
	DeletedAt gorm.DeletedAt `json:"deletedAt"`               // 删除时间
}

type createTable struct {
	table   mysqlTable
	comment string
}

type mysqlTable interface {
	dataFilling()
}

var (
	// 客户端对象
	DB *gorm.DB
	// 初始化错误
	err error
	// 自动创表列表
	tables = make([]createTable, 0)
	// gorm配置
	gormConf = &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	}
)
