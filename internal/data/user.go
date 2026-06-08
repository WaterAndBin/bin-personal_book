package data

import (
	pb "bin-personal-book/api/user/v1"
	"context"

	"bin-personal-book/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) GetList(ctx context.Context, g *pb.LoginParams) (*pb.LoginParams, error) {
	return g, nil
}
