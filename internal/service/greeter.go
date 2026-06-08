package service

import (
	"bin-personal-book/internal/biz"
	"context"
	"fmt"

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

func (s *GreeterService) Login(ctx context.Context, req *pb.LoginParams) (*pb.CommonReply, error) {
	err := s.user.Login(ctx, req)

	fmt.Print(err)

	return &pb.CommonReply{
		Code:    "200",
		Message: "登录成功",
	}, nil
}
