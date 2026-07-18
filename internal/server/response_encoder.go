// internal/server/response_encoder.go
package server

import (
	"bin-personal-book/internal/core"
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	httpTransport "github.com/go-kratos/kratos/v2/transport/http"
)

// 自定义 HTTP 响应编码器
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// protobuf 数据先转 json
	dataJSON, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(v.(proto.Message))

	if err != nil {
		return err
	}

	resp := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    json.RawMessage(dataJSON),
	}

	data, err := json.Marshal(resp)
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
		Reason:  e.Reason,
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
