package service

import (
	"config_tools/app/errors"
	"sync"
)

var (
	lock = &sync.Mutex{}
	// 方法锁
	locks = map[string]*sync.Mutex{}
)

type JsonResponseData struct {
	Code    errors.ErrorCode `json:"code"`
	Message string           `json:"message"`
	Time    int64            `json:"time"`
	Data    interface{}      `json:"data,omitempty"`
}
