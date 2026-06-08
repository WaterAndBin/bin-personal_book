package biz

import (
	pb "bin-personal-book/api/user/v1"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo 方法
type UserRepo interface {
	GetList(ctx context.Context, params *pb.LoginParams) (*pb.LoginParams, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, params *pb.LoginParams) error {
	fmt.Println(params)

	return nil
}
