package data

import (
	"bin-personal-book/internal/biz"
	"bin-personal-book/internal/conf"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type uploadData struct {
	confData *conf.Data
	data     *Data
	log      *log.Helper
}

func NewUploadData(confData *conf.Data, data *Data, logger log.Logger) biz.UploadZip {
	return &uploadData{
		confData: confData,
		data:     data,
		log:      log.NewHelper(logger),
	}
}

// 保存到本地
func (up *uploadData) SaveFile(file multipart.File, header *multipart.FileHeader, name string) (*string, error) {
	dir := filepath.Join(up.confData.Upload.Path, name)

	// 创建文件夹
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// 获取文件后缀
	extension := path.Ext(header.Filename)
	// 生成唯一id的文件名
	newFilename := uuid.New().String() + extension

	dstPath := filepath.Join(dir, newFilename)

	// 创建文件
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// 将上传的文件流写入到目标文件中
	_, err = io.Copy(dst, file)
	if err != nil {
		return nil, err
	}

	// 拿到完整的保存地址
	saveUrl := filepath.Join(
		name,
		newFilename,
	)

	return &saveUrl, nil
}
