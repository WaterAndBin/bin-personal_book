package service

import (
	"bin-personal-book/internal/biz"
	"context"

	pb "bin-personal-book/api/user/v1"
)

type UserService struct {
	pb.UnimplementedGreeterServer

	user *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{
		user: user,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginParams) (*pb.LoginResult, error) {
	return s.user.Login(ctx, req)
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterParams) (*pb.RegisterResult, error) {
	return s.user.Register(ctx, req)
}
