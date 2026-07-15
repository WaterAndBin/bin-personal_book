// internal/server/response_encoder.go
package server

import (
	"bin-personal-book/internal/core"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"

	httpTransport "github.com/go-kratos/kratos/v2/transport/http"
)

// 自定义 HTTP 响应编码器
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	resp := core.Response{
		Code:    200,
		Message: "success",
		Data:    v,
	}

	codec, _ := httpTransport.CodecForRequest(r, "Accept")

	// 将 Response 结构体序列化为字节数组
	data, err := codec.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	return err
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// 转换为 Kratos Error，方便获取 HTTP Code、Message 等信息
	e := errors.FromError(err)

	// 获取最终返回给前端的业务错误码：
	// 优先返回自定义业务码，没有则返回 HTTP 状态码
	code := core.GetErrorCode(err)

	resp := core.Response{
		Code:    code,
		Message: e.Message,
	}

	codec, _ := httpTransport.CodecForRequest(r, "Accept")

	data, _ := codec.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")

	// HTTP 状态码保持 Kratos 原始状态码，
	// 例如 401、404、500，方便网关、浏览器和日志系统识别。
	w.WriteHeader(int(e.Code))

	_, _ = w.Write(data)
}
