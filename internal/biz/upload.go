package biz

import (
	upload "bin-personal-book/api/upload/v1"
	"bin-personal-book/internal/conf"
	"fmt"
	"mime/multipart"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UploadZip interface {
	SaveFile(file multipart.File, header *multipart.FileHeader, name string) (*string, error)
}

type UploadUsecase struct {
	confData *conf.Data
	repo     UploadZip
	log      *log.Helper
}

func NewUploadBiz(confData *conf.Data, repo UploadZip, logger log.Logger) *UploadUsecase {
	return &UploadUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}

func (up *UploadUsecase) Upload(ctx http.Context) (*upload.UploadResult, error) {
	request := ctx.Request()

	// 限制请求体大小并解析表单
	if err := request.ParseMultipartForm(1 << 20); err != nil {
		return nil, errors.BadRequest("error", "文件解析失败或过大")
	}

	if len(request.MultipartForm.File["file"]) == 0 {
		return nil, errors.BadRequest("error", "请上传照片")
	}

	if len(request.MultipartForm.File["file"]) > 1 {
		return nil, errors.BadRequest("error", "只能上传一张图片")
	}

	name := request.FormValue("name")

	if name == "" {
		return nil, errors.BadRequest("error", "缺少name")
	}

	// 拿到文件
	file, header, err := request.FormFile("file")
	if err != nil {
		return nil, err
	}
	// 关闭文件缓存
	defer file.Close()

	if header.Size > 1<<20 {
		return nil, fmt.Errorf("文件大小不能超过 1MB")
	}

	url, saveErr := up.repo.SaveFile(file, header, name)
	if saveErr != nil {
		return nil, saveErr
	}

	return &upload.UploadResult{
		URL: *url,
	}, nil
}
