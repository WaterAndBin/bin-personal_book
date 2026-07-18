package server

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterFileServiceHTTPServer(srv *http.Server) {
	// 文件上传相关的接口
	route := srv.Route("/")
	route.POST("/upload", upload())
}

func upload() http.HandlerFunc {
	return func(ctx http.Context) error {
		req := ctx.Request()

		// 解析 multipart/form-data 请求
		// ParseMultipartForm需在最前面校验，放到后面就会使用默认32MB
		err := req.ParseMultipartForm(1 << 10) // 限制上传文件大小为 10MB
		if err != nil {
			return errors.BadRequest(
				"name_required",
				"体积过大",
			)
		}

		name := req.FormValue("name")

		if name == "" {
			return errors.BadRequest(
				"error",
				"缺少必要参数",
			)
		}

		// 拿到表单文件数据
		file, header, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return err
		}

		defer file.Close()

		fmt.Println("filename:", header.Filename)

		return nil

	}
}
