package biz

import (
	"bin-personal-book/internal/conf"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/log"
)

type UploadUsecase struct {
	confData *conf.Data
	log      *log.Helper
}

func NewUploadBiz(confData *conf.Data, logger log.Logger) *UploadUsecase {
	return &UploadUsecase{confData: confData, log: log.NewHelper(logger)}
}

// 保存到本地
func (up *UploadUsecase) SaveFile(file multipart.File, name string) error {
	dir := filepath.Join("test123123", name)

	fmt.Println(dir)

	err := os.Mkdir(dir, 0755)

	if err != nil {
		return err
	}

	dstPath := filepath.Join(dir, "image")

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
