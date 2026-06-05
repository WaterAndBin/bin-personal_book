package biz

import (
	"context"

	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo 方法
type UserRepo interface {
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, name string) error {
	uc.log.WithContext(ctx).Infof("Login: %v", name)

	if name == "" {
		return errors.New("用户名不能为空")
	}

	return nil
}
