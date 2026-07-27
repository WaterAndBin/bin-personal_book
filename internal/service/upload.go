package service

import (
	"bin-personal-book/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func (s *UploadService) RegisterFileServiceHTTPServer(srv *http.Server) {
	// 文件上传相关的接口
	route := srv.Route("/")
	route.POST("/upload", s.Upload)
}

type UploadService struct {
	up *biz.UploadUsecase
}

func NewUploadService(up *biz.UploadUsecase) *UploadService {
	return &UploadService{
		up: up,
	}
}

func (s *UploadService) Upload(httpCtx http.Context) error {
	h := httpCtx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.up.Upload(httpCtx)
	})
	return httpCtx.Returns(h(httpCtx, nil))
}
