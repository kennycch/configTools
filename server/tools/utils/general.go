package utils

import (
	"strconv"

	"github.com/kennycch/gotools/general"
)

func StrToNumber[T general.Number](value string) T {
	num, _ := strconv.ParseFloat(value, 64)
	return T(num)
}
