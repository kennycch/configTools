package service

import "config_tools/app/errors"

type JsonResponseData struct {
	Code    errors.ErrorCode `json:"code"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data,omitempty"`
	Time    int64            `json:"time"`
}
