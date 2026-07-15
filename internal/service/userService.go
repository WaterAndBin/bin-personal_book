package service

import (
	"bin-personal-book/internal/biz"
	"context"

	pb "bin-personal-book/api/user/v1"
)

func NewUserService(user *biz.UserUsecase) *MainService {
	return &MainService{
		user: user,
	}
}

func (s *MainService) Login(ctx context.Context, req *pb.LoginParams) (*pb.LoginResult, error) {
	return s.user.Login(ctx, req)
}

func (s *MainService) Register(ctx context.Context, req *pb.RegisterParams) (*pb.RegisterResult, error) {
	return s.user.Register(ctx, req)
}
