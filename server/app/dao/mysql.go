package dao

import (
	"config_tools/config"
	"config_tools/tools/log"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	// gorm配置
	gormConf := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	}
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

func autoCreateTable() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8mb4").AutoMigrate(dst...)
}
