package biz

import (
	"bin-personal-book/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo 方法
type TagsRepo interface {
}

type TagsUsecase struct {
	confData *conf.Data
	repo     UserRepo
	log      *log.Helper
}

func NewTagsUseBiz(confData *conf.Data, repo UserRepo, logger log.Logger) *TagsUsecase {
	return &TagsUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}
