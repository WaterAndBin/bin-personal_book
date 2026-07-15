package service

import (
	pb "bin-personal-book/api/user/v1"
	"bin-personal-book/internal/biz"

	"github.com/google/wire"
)

type MainService struct {
	pb.UnimplementedGreeterServer

	user *biz.UserUsecase
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)
