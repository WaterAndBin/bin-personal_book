package service

import (
	"bin-personal-book/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"

	pb "bin-personal-book/api/tags/v1"
)

func (s *TagsService) RegisterTagsServiceHTTPServer(srv *http.Server) {
	// 文件上传相关的接口
	route := srv.Route("/")
	route.POST("/getBillTagsList", s.GetBillTagsList)
}

type TagsService struct {
	pb.UnimplementedGreeterServer

	tags *biz.TagsUsecase
}

func NewTagsService(tag *biz.TagsUsecase) *TagsService {
	return &TagsService{
		tags: tag,
	}
}

func (s *TagsService) GetBillTagsList(httpCtx http.Context) error {
	h := httpCtx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		params := req.(*pb.GetBillTagsListParams)

		return s.tags.GetBillTagsList(ctx, params)
	})
	return httpCtx.Returns(h(httpCtx, &pb.GetBillTagsListParams{}))
}

func (s *TagsService) UpdateBillTags(ctx context.Context, req *pb.UpdateBillTagsParams) (*pb.UpdateBillTagsResult, error) {
	return s.tags.UpdateBillTags(ctx, req)
}
