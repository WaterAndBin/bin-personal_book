package service

import (
	"bin-personal-book/internal/biz"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UploadService struct {
	up *biz.UploadUsecase
}

func NewUploadService(up *biz.UploadUsecase) *UploadService {
	return &UploadService{
		up: up,
	}
}

func (s *UploadService) Upload(ctx http.Context) error {
	req := ctx.Request()

	// 限制请求体大小并解析表单
	if err := req.ParseMultipartForm(1 << 20); err != nil {
		return errors.BadRequest("error", "文件解析失败或过大")
	}

	if len(req.MultipartForm.File["file"]) != 1 {
		return errors.BadRequest("error", "只能上传一张图片")
	}

	name := req.FormValue("name")

	if name == "" {
		return errors.BadRequest("error", "缺少name")
	}

	// 拿到文件
	file, header, err := req.FormFile("file")
	if err != nil {
		return err
	}
	// 关闭文件缓存
	defer file.Close()

	if header.Size > 1<<20 {
		return fmt.Errorf("文件大小不能超过 1MB")
	}

	s.up.SaveFile(file, name)

	return nil
}
