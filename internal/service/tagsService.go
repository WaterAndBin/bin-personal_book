package service

import (
	"bin-personal-book/internal/biz"
	"context"

	pb "bin-personal-book/api/tags/v1"
)

type TagsService struct {
	pb.UnimplementedGreeterServer

	tags *biz.TagsUsecase
}

func NewTagsService(tag *biz.TagsUsecase) *TagsService {
	return &TagsService{
		tags: tag,
	}
}

func (s *TagsService) GetBillTagsList(ctx context.Context, req *pb.GetBillTagsListParams) (*pb.GetBillTagsListResult, error) {
	return s.tags.GetBillTagsList(ctx, req)
}
