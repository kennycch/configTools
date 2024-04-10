package dao

import "config_tools/tools/lifecycle"

type MysqlRegister struct{}

func (m *MysqlRegister) Start() {
	// 初始化
	mysqlInit()
	// 自动创表
	autoCreateTable()
	// 数据填充
	dataFilling()
}

func (m *MysqlRegister) Priority() uint32 {
	return lifecycle.NormalPriority + 10000
}

func (m *MysqlRegister) Stop() {
	mysqlClose()
}

func NewMysqlRegister() *MysqlRegister {
	return &MysqlRegister{}
}
