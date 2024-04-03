package config

import (
	"config_tools/tools/lifecycle"
	"flag"

	"github.com/go-ini/ini"
)

type ConfigRegister struct{}

func (c *ConfigRegister) Start() {
	flag.Parse()
	// 读取env配置文件
	if env, err := ini.Load(*filePath); err != nil {
		panic(err)
	} else {
		MapConfig(env)
	}
}

func (c *ConfigRegister) Priority() uint32 {
	return lifecycle.HighPriority + 10000
}

func (c *ConfigRegister) Stop() {

}

func NewConfigRegister() *ConfigRegister {
	return &ConfigRegister{}
}
