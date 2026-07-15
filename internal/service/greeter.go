package service

import (
	"bin-personal-book/internal/biz"
	"context"

	pb "bin-personal-book/api/user/v1"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	user *biz.UserUsecase
}

func NewGreeterService(user *biz.UserUsecase) *GreeterService {
	return &GreeterService{
		user: user,
	}
}

func (s *GreeterService) Login(ctx context.Context, req *pb.LoginParams) (*pb.LoginResult, error) {
	_, err := s.user.Login(ctx, req)

	return &pb.LoginResult{
		Token: "123",
	}, err
}
