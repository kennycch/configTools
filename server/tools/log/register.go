package log

import (
	"config_tools/config"
	"config_tools/tools/lifecycle"

	"github.com/kennycch/gotools/log"
)

type LogRegister struct{}

func (c *LogRegister) Start() {
	logLevel := log.InfoLevel
	if config.App.Debug {
		logLevel = log.DebugLevel
	}
	log.InitLog(
		config.Log.LogPath,
		config.App.AppName,
		config.Log.LogDay,
		logLevel,
	)
	isInit = true
}

func (c *LogRegister) Priority() uint32 {
	return lifecycle.HighPriority + 8000
}

func (c *LogRegister) Stop() {

}

func NewLogRegister() *LogRegister {
	return &LogRegister{}
}
