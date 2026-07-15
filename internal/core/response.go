package core

import (
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
)

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	// 参数错误
	ParamError int32 = 401
	// Token 已过期
	TokenExpired int32 = 402
	// 无权限
	Forbidden int32 = 403
	// 系统异常
	ServerError int32 = 500
)

// NewError 创建业务异常。
// 不传 code 时，默认使用 ServerError。
// 传入 code 时，使用指定的业务错误码。
func NewError(message string, codes ...int32) error {
	// 默认业务错误码
	code := ServerError

	// 如果调用方传入了业务错误码，则使用传入的错误码
	if len(codes) > 0 {
		code = codes[0]
	}

	// 创建 Kratos Error（HTTP 状态码默认为 500）
	err := errors.InternalServer("BUSINESS", message)

	// 将业务错误码存入 Metadata，供 ErrorEncoder 统一解析
	err.Metadata = map[string]string{
		"code": strconv.Itoa(int(code)),
	}

	return err
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
