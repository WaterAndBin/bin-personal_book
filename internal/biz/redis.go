package biz

import (
	"bin-personal-book/internal/conf"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type RedisZip interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type RedisUsecase struct {
	confData *conf.Data
	repo     RedisZip
	log      *log.Helper
}

func NewRedisUseBiz(confData *conf.Data, repo RedisZip, logger log.Logger) *RedisUsecase {
	return &RedisUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}

func (uc *RedisUsecase) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) error {
	return uc.repo.Set(
		ctx,
		key,
		value,
		expiration,
	)
}

func (uc *RedisUsecase) Get(ctx context.Context, key string) (string, error) {
	return uc.repo.Get(
		ctx,
		key,
	)
}
