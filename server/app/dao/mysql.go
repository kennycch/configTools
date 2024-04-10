package dao

import (
	"config_tools/config"
	"config_tools/tools/log"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysqlInit 初始化Mysql
func mysqlInit() {
	// 链接设置
	dia := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Account,
		config.Mysql.Password,
		config.Mysql.Addr,
		config.Mysql.DataBase,
	)
	// 尝试连接
	DB, err = gorm.Open(mysql.Open(dia), gormConf)
	if err != nil {
		panic(err)
	}
	sqlDb, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	log.Info("mysql init done ", zap.String("conn", dia))
}

// mysqlClose 关闭客户端
func mysqlClose() {
	if DB != nil {
		sqlDb, err := DB.DB()
		if err != nil {
			log.Error("mysql close err", zap.Error(err))
			return
		}
		sqlDb.Close()
	}
}

// 自动创表
func autoCreateTable() {
	for _, table := range tables {
		DB.Set("gorm:table_options", fmt.Sprintf("ENGINE=InnoDB,CHARSET=utf8mb4,COMMENT='%s'", table.comment)).AutoMigrate(table.table)
	}
}

// 数据填充
func dataFilling() {
	for _, table := range tables {
		table.table.dataFilling()
	}
}

func (m *Model) dataFilling() {

}

// GetOffset 计算偏移量
func GetOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}
