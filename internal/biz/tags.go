package biz

import (
	tags "bin-personal-book/api/tags/v1"
	"bin-personal-book/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsZip interface {
	GetBillTagsList(ctx context.Context, params *tags.GetBillTagsListParams) (*tags.GetBillTagsListResult, error)
	UpdateBillTags(ctx context.Context, params *tags.UpdateBillTagsParams) (*tags.UpdateBillTagsResult, error)
}

type TagsUsecase struct {
	confData *conf.Data
	repo     TagsZip
	log      *log.Helper
}

func NewTagsUseBiz(confData *conf.Data, repo TagsZip, logger log.Logger) *TagsUsecase {
	return &TagsUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}

func (uc *TagsUsecase) GetBillTagsList(ctx context.Context, params *tags.GetBillTagsListParams) (*tags.GetBillTagsListResult, error) {
	return uc.repo.GetBillTagsList(ctx, params)
}

func (uc *TagsUsecase) UpdateBillTags(ctx context.Context, params *tags.UpdateBillTagsParams) (*tags.UpdateBillTagsResult, error) {
	return uc.repo.UpdateBillTags(ctx, params)
}
