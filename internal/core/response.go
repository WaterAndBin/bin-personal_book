package core

import (
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
)

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
	Data    interface{} `json:"data,omitempty"`
}

// GetErrorCode 获取业务错误码。
// 如果错误中没有业务错误码，则返回 HTTP 状态码。
func GetErrorCode(err error) int32 {
	e := errors.FromError(err)

	if value, ok := e.Metadata["code"]; ok {
		if code, err := strconv.Atoi(value); err == nil {
			return int32(code)
		}
	}

	return e.Code
}
