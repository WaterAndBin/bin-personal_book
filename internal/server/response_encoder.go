// internal/server/response_encoder.go
package server

import (
	"bin-personal-book/internal/core"
	"net/http"

	httpTransport "github.com/go-kratos/kratos/v2/transport/http"
)

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	resp := core.Response{
		Code:    0,
		Message: "success",
		Data:    v,
	}

	codec, _ := httpTransport.CodecForRequest(r, "Accept")

	data, err := codec.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	return err
}
