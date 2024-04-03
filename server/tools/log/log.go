package log

import (
	"github.com/kennycch/gotools/log"
	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	if !isInit {
		return
	}
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	if !isInit {
		return
	}
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	if !isInit {
		return
	}
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if !isInit {
		return
	}
	log.Error(msg, fields...)
}
