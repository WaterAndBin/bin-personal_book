package service

import (
	"bin-personal-book/internal/biz"
	"context"

	pb "bin-personal-book/api/user/v1"
)

func NewTagsService(tag *biz.TagsUsecase) *MainService {
	return &MainService{
		tag: tag,
	}
}

func (s *MainService) GetBillTagsList(ctx context.Context, req *pb.LoginParams) (*pb.LoginResult, error) {
	return s.user.Login(ctx, req)
}
