package server

import (
	"encoding/json"
	nethttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/go-kratos/kratos/v2/transport/http"

	jwtV5 "github.com/golang-jwt/jwt/v5"
)

// 自定义 HTTP 响应编码器
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	var dataJSON []byte
	var err error

	switch data := v.(type) {

	// 兼容多种上传回调
	case proto.Message:
		// protobuf 类型
		dataJSON, err = protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true,
		}.Marshal(data)
	default:
		// 普通 struct / map / string
		dataJSON, err = json.Marshal(data)
	}

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
	w.WriteHeader(nethttp.StatusOK)

	_, err = w.Write(data)

	return err
}

func ErrorEncoder(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	se := errors.FromError(err)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		int(se.Code),
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"code":    se.Code,
			"reason":  se.Reason,
			"message": se.Message,
		},
	)

}

// 手动校验jwt是否过期
func JWTHandler(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(ctx http.Context) error {

		req := ctx.Request()

		auth := req.Header.Get("Authorization")
		if auth == "" {
			return errors.Unauthorized(
				"UNAUTHORIZED",
				"JWT token is missing",
			)
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		_, err := jwtV5.Parse(tokenString, func(token *jwtV5.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil {
			return errors.Unauthorized(
				"UNAUTHORIZED",
				"JWT token invalid",
			)
		}

		return next(ctx)
	}
}
